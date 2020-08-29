# CEasy-Backend

## NPE가 난무하는 안전행정부 재난 문자 구조 정리

### 가장 최근의 ID가져오기

post: http://www.safekorea.go.kr/idsiSFK/bbs/user/selectBbsList.do  
request:

```json
{
  "bbs_searchInfo": {
    "bbs_no": "63",
    "pageUnit": "1"
  }
}
```

response:

```json
{
  "rtnResult": {
    "resultCode": "성공: 0, 실패: -1",
    "resultMsg": "성공여부 관련 메시지"
  },
  "bbsList": [
    {
      "FRST_REGIST_DT": "작성일",
      "BBS_ORDR": "ID(int)",
      "SJ": "제목"
    }
  ]
}
```

### 세부정보 가져오기

POST: https://www.safekorea.go.kr/idsiSFK/bbs/user/selectBbsView.do

```json
{
  "bbs_searchInfo": {
    "bbs_no": "63",
    "bbs_ordr": "가져올 ID"
  }
}
```

response:

```json
{
  "rtnResult": {
    "resultCode": "성공: 0, 실패: -1",
    "resultMsg": "성공여부 관련 메시지"
  },
  "bbsMap": {
    "sj": "제목",
    "cn": "내용",
    "frst_regist_dt": "글 작성일"
  }
}
```
