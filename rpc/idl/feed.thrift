namespace go feed

typedef i64 int64

struct Video {
    1: int64 Id;
    2: int64 Author_id;
    3: string Play_url;
    4: string Cover_url;
    5: string Upload_time;
    6: string Title;
    7: int64 Favorite_count;
    8: int64 Comment_count;
}

struct FeedRequest {
    1: optional string Laest_time;
    2: optional int64 Author_id;
}

struct FeedResponse {
    1: int64 Status_code;
    2: list<Video> Videos_list;
}

service FeedService {
    FeedResponse GetVideo(1: FeedRequest req);
}

