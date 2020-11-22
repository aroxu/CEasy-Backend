# CEasy-Backend

## NPE가 난무하는 행정안전부 재난 문자 구조 정리

### SafeKorea 사이트에서 가장 최근의 ID가져오기

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

### SafeKorea 사이트에서 세부정보 가져오기

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

## 직접 서버 돌리기

1. MySQL 8.0을 설치 후 세팅.
2. `ceasy` 라는 이름을 가진 데이터베이스를 생성. 테이블 생성과 구조 설정은 불필요. 런타임에서 매 실행시마다 AutoMigrate를 수행하기 때문.
3. `ceasy.config.inc.json` 파일을 `ceasy.config.json`으로 파일 이름을 변경한 뒤, 안의 내용을 적절히 기입한다.
4. `go build && ./CEasy-Backend` 실행. 또는 `go run main.go`를 실행

## **request / response 정리**

`code` : 성공 여부  
`count` : 서버가 현제 가지고 있는 데이터의 양을 표시  
`data` : request에 맞게 산출된 데이터, List 형태로 표시됨  
┣`Date` : 재난 문자 발송일 및 시간 기록  
┣`ID` : 재난 문자 발행 ID  
┣`area` : 발송 주최  
┣`area_detail` : 조금 더 자세한 발송 주최  
┗`content` : 재난 문자의 내용  
`message` : 오류 메세지

#### 아무것도 붙이지 않았을때

ex) GET : localhost:9096/api/cbs/v0/

response:

```json
{
  "code": "SUCCESS",
  "count": 2512,
  "data": [
    {
      "Date": "2020-09-11T17:47:19Z",
      "ID": 56281,
      "area": "경남도청",
      "area_detail": "경상남도 전체",
      "content": "8.30.~9.10. 함양군 지리산택시를 타신 분은 즉시 전화해 코로나19 상담 바랍니다.(055-960-8014~15,8024,8031,8033)"
    },
    {
      "Date": "2020-09-11T17:39:57Z",
      "ID": 56280,
      "area": "포항시청",
      "area_detail": "경상북도 포항시",
      "content": "9.2.(수)14~18시 칠곡군 동명면 평산아카데미 장뇌삼 사업설명회(코로나19 확진자발생)에 참석자는 즉시 인근 보건소에서 검사 받으시기 바랍니다."
    },
    {
      "Date": "2020-09-11T17:39:57Z",
      "ID": 56279,
      "area": "안동시청",
      "area_detail": "경상북도 안동시",
      "content": "최근 타지역 확진자발생관련/방문판매업 관련종사자는 규모 관계없이 사업설명회 개최 및 참석을 자제하여 주시기 바랍니다."
    },
    ...
  ],
  "message": ""
}
```

#### 여태까지 발신 된 위치를 조회할때

ex) GET : localhost:9096/api/cbs/v0/location?value=서

> 발신 기록중에 '서'라는 지역에서 발신한 데이터를 산출하여 위치만 반환

response:

```json
{
  "code": "SUCCESS",
  "data": [
    "강서구청",
    "서구청",
    "서대문구",
    "서대문구청",
    "서산시",
    "서산시청",
    "서울시청",
    "서천군청",
    "서초구청",
    "인천서구청"
  ],
  "message": ""
}
```

#### 위치기반으로 검색할때 (대략적)

ex) GET : localhost:9096/api/cbs/v0/?area=인천시

> 주소(area 필드)에 '인천시'라는 글자가 들어가는 모든 정보 반환

response:

```json
{
  "code": "SUCCESS",
  "count": 3,
  "data": [
    {
      "Date": "2020-09-11T13:51:50Z",
      "ID": 56200,
      "area": "인천시",
      "area_detail": "인천광역시 전체",
      "content": "추석연휴9.30~10.4인천가족공원미운영/미리성묘,온라인성묘 http://grave.insiseol.or.kr적극이용바랍니다코로나우울상담1577-0199"
    },
    {
      "Date": "2020-09-06T20:00:31Z",
      "ID": 54731,
      "area": "인천시",
      "area_detail": "인천광역시 전체",
      "content": "9월7일태풍영향권.낙하물주의,해안가접근금지,외출자제주의바랍니다.시민행동요령링크:https://www.incheon.go.kr/safe/SAFE020101"
    },
    {
      "Date": "2020-09-06T12:07:21Z",
      "ID": 54456,
      "area": "인천시",
      "area_detail": "인천광역시 전체",
      "content": "강화된 사회적 거리두기 2단계 연장시행. 3밀(밀폐&middot;밀집&middot;밀접)환경을 피하고, 감염의심자(발열, 마른기침,피로감 등)는 주소지 보건소에서 검사바랍니다."
    }
  ],
  "message": ""
}
```

#### 위치기반으로 검색할때 (조금 더 세부적)

ex) GET : localhost:9096/api/cbs/v0/?area_detail=서대문

> 상세 주소(area_detail 필드)에 '서대문'라는 글자가 들어가는 모든 정보 반환

response:

```json
{
  "code": "SUCCESS",
  "count": 11,
  "data": [
    {
      "Date": "2020-09-11T14:57:21Z",
      "ID": 56219,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 112~113번째 확진자 발생. 역학조사 중. 자세한 사항 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-10T16:16:31Z",
      "ID": 56011,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 109~111번째 확진자 발생. 역학조사 중. 자세한 사항 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-09T17:32:50Z",
      "ID": 55823,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 107~108번째 확진자 발생. 역학조사 중. 자세한 사항 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-08T15:58:44Z",
      "ID": 55560,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 105~106번째 확진자 발생. 역학조사 중. 자세한 사항 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-07T14:32:31Z",
      "ID": 55327,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 104번째 확진자 발생. 역학조사 중. 자세한 사항 추후 구홈페이지, 블로그 공개 예정입니다.blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-07T13:49:50Z",
      "ID": 55312,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "태풍 「하이선」 금일 오후 2~3시 최근접 예상되니 산림(안산, 북한산 등) 및 하천 출입을 자제하고 간판 등 낙하물 주의하시기 바랍니다."
    },
    {
      "Date": "2020-09-06T12:18:30Z",
      "ID": 54463,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 103번째 확진자 발생. 역학조사중. 자세한 사항 추후 구홈페이지, 블로그 공개 예정입니다.blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-05T17:33:50Z",
      "ID": 54288,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 102번째 확진자 발생. 역학조사중. 자세한 사항 추후 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-05T13:44:13Z",
      "ID": 54191,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 101번째 확진자 발생. 역학조사 중. 자세한 사항 구홈페이지, 블로그 공개 예정입니다blog.naver.com/sdmstory"
    },
    {
      "Date": "2020-09-04T17:54:50Z",
      "ID": 53945,
      "area": "서대문구",
      "area_detail": "서울특별시 서대문구",
      "content": "태풍 ‘하이선’ 관련, 서대문구 보건소 선별진료소가 9.7.(월) 오전 09:00~13:00 단축운영하오니 방문 시 차질 없으시길 바랍니다."
    },
    {
      "Date": "2020-09-04T14:58:30Z",
      "ID": 53874,
      "area": "서대문구청",
      "area_detail": "서울특별시 서대문구",
      "content": "코로나19. 98~100번째 확진자 발생. 역학조사 진행 중. 자세한 사항 추후 구홈페이지 및 블로그 공개 예정입니다."
    }
  ],
  "message": ""
}
```

#### 한도 조절

ex) GET : localhost:9096/api/cbs/v0/limit=5

> 산출된 데이터에서 가장 최근 5개의 항목만 산출

response:

```json
{
  "code": "SUCCESS",
  "count": 2523,
  "data": [
    {
      "Date": "2020-09-11T18:08:51Z",
      "ID": 56292,
      "area": "포천시청",
      "area_detail": "경기도 포천시",
      "content": "포천시 코로나19 62번~64번 확진자 발생(영북면 거주 3명) 역학조사 중 동선은 홈페이지 및 블로그(c11.kr/hupu)참조하시기 바랍니다."
    },
    {
      "Date": "2020-09-11T18:05:59Z",
      "ID": 56291,
      "area": "고양시청",
      "area_detail": "경기도 고양시",
      "content": "코로나19 확진자 1명 발생 ▶313번(화전동 거주) [홈페이지][카카오톡 고양시 채널] https://han.gl/iMmgw 참조 바랍니다."
    },
    {
      "Date": "2020-09-11T18:05:31Z",
      "ID": 56290,
      "area": "울릉군청",
      "area_detail": "경상북도 울릉군",
      "content": "태풍 응급복구 작업으로 별도 해제시까지 통구미터널~남통터널~남양터널 구간의 통행을 금지하오니 우회도로를 이용하시기 바랍니다."
    },
    {
      "Date": "2020-09-11T18:04:18Z",
      "ID": 56289,
      "area": "광주광역시",
      "area_detail": "광주광역시 전체",
      "content": "오늘 현재 확진자 3명 발생하여 총 475명입니다. 473, 474번은 남구 주월동, 475번 북구 용두동 거주자로 동선은 조사중입니다."
    },
    {
      "Date": "2020-09-11T18:04:19Z",
      "ID": 56288,
      "area": "순천시청",
      "area_detail": "전라남도 순천시",
      "content": "빠른 일상 복귀 위하여 주말 타지역 방문 삼가, 집에만 머물기, 아프면 쉬기, 마스크 착용 생활화 등 생활방역 수칙을 지켜주시기 바랍니다."
    }
  ],
  "message": ""
}
```

#### 날짜 구간 지정

ex) GET : localhost:9096/api/cbs/v0/?start=2020-09-11%2018:08:01

> 산출된 데이터에서 가장 최근 5개의 항목만 산출, 시간 형식은 yyyy-mm-dd hh:mm:ss로 지정, 동일하게 end 속성도 적용

response:

```json
{
  "code": "SUCCESS",
  "count": 5,
  "data": [
    {
      "Date": "2020-09-11T18:14:31Z",
      "ID": 56296,
      "area": "남동구청",
      "area_detail": "인천광역시 남동구",
      "content": "코로나19 추가 확진자 1명 발생(간석1동, 확진자 접촉자) 세부내용은 홈페이지 및 블로그에 게시하였습니다. reurl.kr/38FA34A0WX"
    },
    {
      "Date": "2020-09-11T18:14:19Z",
      "ID": 56295,
      "area": "동대문구청",
      "area_detail": "서울특별시 동대문구",
      "content": "오늘 우리 동대문구 확진자 수가 다시 3명으로 증가했습니다. 구민 여러분께서는 기본 예방수칙을 준수해 주시기를 각별히 부탁드립니다."
    },
    {
      "Date": "2020-09-11T18:13:57Z",
      "ID": 56294,
      "area": "동대문구청",
      "area_detail": "서울특별시 동대문구",
      "content": "134번(용신동,남),135번(전농1동,여),136번(장안2동,남) 확진자 발생, 자세한 사항은 ddm4you.blog.me 참조하시기 바랍니다."
    },
    {
      "Date": "2020-09-11T18:08:50Z",
      "ID": 56293,
      "area": "칠곡군청",
      "area_detail": "경상북도 칠곡군",
      "content": "경주 67번 확진자 동명면 장뇌삼 사업설명회 방문 관련 참석자 명단 전원 확보. 현재 역학 조사 중임을 알려드립니다."
    },
    {
      "Date": "2020-09-11T18:08:51Z",
      "ID": 56292,
      "area": "포천시청",
      "area_detail": "경기도 포천시",
      "content": "포천시 코로나19 62번~64번 확진자 발생(영북면 거주 3명) 역학조사 중 동선은 홈페이지 및 블로그(c11.kr/hupu)참조하시기 바랍니다."
    }
  ],
  "message": ""
}
```

#### 여러개의 값 적용

위의 limit, area, area_detail, start, end 등의 속성은 중첩하여 사용할 수 있다.
단, start -> end -> area -> area_detail -> limit 의 순서대로 옵션을 주어야 제대로 결과가 나온다.

#### 오프셋 지정

ex) GET : localhost:9096/api/cbs/v0/?limit=5&offset=20

> 조건에 맞게 산출된 데이터에서 20개 뒤의 데이터부터 지정

response:

```json
{
  "code": "SUCCESS",
  "count": 46105,
  "data": [
    {
      "Date": "2020-09-11T18:14:31Z",
      "ID": 56296,
      "area": "남동구청",
      "area_detail": "인천광역시 남동구",
      "content": "코로나19 추가 확진자 1명 발생(간석1동, 확진자 접촉자) 세부내용은 홈페이지 및 블로그에 게시하였습니다. reurl.kr/38FA34A0WX"
    },
    {
      "Date": "2020-09-11T18:14:19Z",
      "ID": 56295,
      "area": "동대문구청",
      "area_detail": "서울특별시 동대문구",
      "content": "오늘 우리 동대문구 확진자 수가 다시 3명으로 증가했습니다. 구민 여러분께서는 기본 예방수칙을 준수해 주시기를 각별히 부탁드립니다."
    },
    {
      "Date": "2020-09-11T18:13:57Z",
      "ID": 56294,
      "area": "동대문구청",
      "area_detail": "서울특별시 동대문구",
      "content": "134번(용신동,남),135번(전농1동,여),136번(장안2동,남) 확진자 발생, 자세한 사항은 ddm4you.blog.me 참조하시기 바랍니다."
    },
    {
      "Date": "2020-09-11T18:08:50Z",
      "ID": 56293,
      "area": "칠곡군청",
      "area_detail": "경상북도 칠곡군",
      "content": "경주 67번 확진자 동명면 장뇌삼 사업설명회 방문 관련 참석자 명단 전원 확보. 현재 역학 조사 중임을 알려드립니다."
    },
    {
      "Date": "2020-09-11T18:08:51Z",
      "ID": 56292,
      "area": "포천시청",
      "area_detail": "경기도 포천시",
      "content": "포천시 코로나19 62번~64번 확진자 발생(영북면 거주 3명) 역학조사 중 동선은 홈페이지 및 블로그(c11.kr/hupu)참조하시기 바랍니다."
    }
  ],
  "message": ""
}
```
