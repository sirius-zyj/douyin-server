package mysql2redis

import (
	"time"

	"github.com/duanhf2012/origin/log"
	"github.com/go-mysql-org/go-mysql/replication"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
)

type CommonEventHandler struct {
	//canal.DummyEventHandler
	CurrOutput *Output
}

func (handler *CommonEventHandler) OnRow(e *canal.RowsEvent) error {
	// log.Debug("OnRow, action:%s  %v", e.Action, e.Rows)
	entity := RawLogEntityPoll.Get().(*RawLogEntity)
	entity.Action = e.Action
	entity.Rows = e.Rows
	entity.TableName = e.Table.Name
	entity.Header = []string{}
	entity.HeaderMap = map[string]int{}
	//entity.ValueMap = map[string]interface{}{}

	for columnIndex, currColumn := range e.Table.Columns {
		entity.Header = append(entity.Header, currColumn.Name)
		entity.HeaderMap[currColumn.Name] = columnIndex
	}
	handler.CurrOutput.Write(entity)

	return nil
}

func (handler *CommonEventHandler) String() string {
	return "MyEventHandler"
}

func (handler *CommonEventHandler) OnRotate(rep *replication.EventHeader, e *replication.RotateEvent) error {
	log.Debug("OnRotate: %s,%d", e.NextLogName, uint32(e.Position))

	handler.CurrOutput.WritePos(&mysql.Position{Name: string(e.NextLogName), Pos: uint32(e.Position)})
	return nil
}

func (handler *CommonEventHandler) OnTableChanged(rep *replication.EventHeader, schema string, table string) error {
	return nil
}

func (handler *CommonEventHandler) OnDDL(rep *replication.EventHeader, nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	//log.Debug("OnDDL")
	return nil
}

func (handler *CommonEventHandler) OnXID(*replication.EventHeader, mysql.Position) error {
	//log.Debug("OnXID")

	return nil
}

func (handler *CommonEventHandler) OnGTID(*replication.EventHeader, mysql.GTIDSet) error {
	//log.Debug("OnGTID")

	return nil
}

func (handler *CommonEventHandler) OnPosSynced(rep *replication.EventHeader, pos mysql.Position, set mysql.GTIDSet, force bool) error {

	handler.CurrOutput.WritePos(&mysql.Position{Name: pos.Name, Pos: pos.Pos})
	return nil
}

type CommonCanalMeta struct {
	Name           string
	ConfigFilePath string
	Config         *canal.Config
	Canal          *canal.Canal
	CurrOutput     *Output
}

// func (handler *CommonCanalMeta) RunWithConfig(filePath string, name string, out *output.Output, pos *config.Pos) {
func (meta *CommonCanalMeta) RunWithConfig(name string, conf *CanalConfig) error {

	currOutput, createOutputErr := CreateByName(name)
	if createOutputErr != nil {
		log.Error(createOutputErr.Error())
		return createOutputErr
	}

	meta.Name = name
	meta.CurrOutput = currOutput
	cfg, loadConfigErr := canal.NewConfigWithFile(conf.Cancalconfigpath)

	if loadConfigErr != nil {
		log.Error("load canal's config failed:%s", loadConfigErr.Error())
		return loadConfigErr
	}
	currCanal, createCanalErr := canal.NewCanal(cfg)

	if createCanalErr != nil {
		log.Error("Init canal failed:%s", createCanalErr.Error())
		return createCanalErr
	}

	go meta.CurrOutput.Run()
	meta.Canal = currCanal

	currCanal.SetEventHandler(&CommonEventHandler{CurrOutput: meta.CurrOutput})
	if conf.LogPos != nil {
		startPos := mysql.Position{Name: conf.LogPos.Name, Pos: conf.LogPos.Pos}
		log.Debug("Run with pos: %v %v", conf.LogPos.Name, conf.LogPos.Pos)
		currCanal.RunFrom(startPos)
	} else {
		log.Debug("Run without pos")
		currCanal.Run()
	}

	return nil
}

func (meta *CommonCanalMeta) Close() {

	meta.Canal.Close()
	meta.waitEmpty()
}

func (meta *CommonCanalMeta) waitEmpty() {
	for {
		if meta.empty() {
			break
		}
		time.Sleep(time.Microsecond * 100)
	}
}

func (meta *CommonCanalMeta) empty() bool {
	return meta.CurrOutput.Empty()
}
