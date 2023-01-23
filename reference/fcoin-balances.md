<!-- Generator: Widdershins v4.0.1 -->

<h1 id="fcoin-balances">fcoin-balances v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="http://localhost:8080">http://localhost:8080</a>

<h1 id="fcoin-balances-default">Default</h1>

## get-users-userId

<a id="opIdget-users-userId"></a>

`GET /users/{user_id}`

*独自コイン残高・履歴の取得*

<h3 id="get-users-userid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|user_id|path|integer|true|none|

> Example responses

> Example response

```json
{
  "balance": {
    "id": 1,
    "amount": 4000
  },
  "history": [
    {
      "id": 12345,
      "user_id": 1,
      "balance_id": 1,
      "amount": 5000,
      "processed_at": "2019-08-24T14:15:22Z"
    },
    {
      "id": 12345,
      "user_id": 1,
      "balance_id": 1,
      "transfer": {
        "id": 12345,
        "from_user": 1,
        "from_balance": 1,
        "to_user": 2,
        "to_balance": 2,
        "amount": 1000,
        "processed_at": "2019-08-24T14:15:22Z"
      },
      "amount": -1000,
      "processed_at": "2019-08-24T14:15:22Z"
    }
  ]
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

<h3 id="get-users-userid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Example response|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Example response|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Example response|Inline|

<h3 id="get-users-userid-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» balance|[balance](#schemabalance)|false|none|none|
|»» id|integer|false|none|none|
|»» amount|integer|false|none|none|
|» history|[[balance_trans](#schemabalance_trans)]|false|none|none|
|»» balance_trans|[balance_trans](#schemabalance_trans)|false|none|none|
|»»» id|integer|true|none|none|
|»»» user_id|integer|true|none|none|
|»»» balance_id|integer|true|none|none|
|»»» transfer|[transer_trans](#schematranser_trans)|false|none|none|
|»»»» id|integer|true|none|none|
|»»»» from_user|integer|true|none|none|
|»»»» from_balance|integer|true|none|none|
|»»»» to_user|integer|true|none|none|
|»»»» to_balance|integer|true|none|none|
|»»»» amount|integer|false|none|none|
|»»»» processed_at|string(date-time)|false|none|none|
|»»» amount|integer|true|none|none|
|»»» processed_at|string(date-time)|true|none|none|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **404**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post-users-user_id

<a id="opIdpost-users-user_id"></a>

`POST /users/{user_id}`

*独自コイン残高への追加・消費*

> Body parameter

```json
{
  "amount": 1000
}
```

<h3 id="post-users-user_id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|false|none|
|» amount|body|integer|false|none|
|user_id|path|integer|true|none|

> Example responses

> Example response

```json
{
  "id": 12345,
  "user_id": 1,
  "balance_id": 1,
  "amount": 1000,
  "processed_at": "2019-08-24T14:15:22Z"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

<h3 id="post-users-user_id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|[balance_trans](#schemabalance_trans)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Example response|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Example response|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Example response|Inline|

<h3 id="post-users-user_id-responseschema">Response Schema</h3>

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **404**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post-users-user_id-transfer

<a id="opIdpost-users-user_id-transfer"></a>

`POST /users/{user_id}/transfer`

*独自コインの送金*

> Body parameter

```json
{
  "user_id": 2,
  "amount": 1000
}
```

<h3 id="post-users-user_id-transfer-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|false|none|
|» user_id|body|integer|false|none|
|» amount|body|integer|false|none|
|user_id|path|integer|true|none|

> Example responses

> Example response

```json
{
  "id": 12345,
  "user_id": 1,
  "balance_id": 1,
  "transfer": {
    "id": 67890,
    "from_user": 1,
    "from_balance": 1,
    "to_user": 2,
    "to_balance": 2,
    "amount": 1000,
    "processed_at": "2019-08-24T14:15:22Z"
  },
  "amount": -1000,
  "processed_at": "2019-08-24T14:15:22Z"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

<h3 id="post-users-user_id-transfer-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|[balance_trans](#schemabalance_trans)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Example response|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Example response|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Example response|Inline|

<h3 id="post-users-user_id-transfer-responseschema">Response Schema</h3>

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **404**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post-user

<a id="opIdpost-user"></a>

`POST /user`

*ユーザー登録（独自コイン残高の作成）*

> Body parameter

```json
{
  "name": "taro"
}
```

<h3 id="post-user-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|false|none|
|» name|body|string|true|none|

> Example responses

> Example response

```json
{
  "user": {
    "id": 1,
    "name": "taro"
  },
  "balance": {
    "id": 1,
    "amount": 0
  }
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

```json
{
  "err_code": "string",
  "message": "string"
}
```

<h3 id="post-user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Example response|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Example response|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Example response|Inline|

<h3 id="post-user-responseschema">Response Schema</h3>

Status Code **201**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» user|[user](#schemauser)|false|none|none|
|»» id|integer|true|none|none|
|»» name|string|true|none|none|
|» balance|[balance](#schemabalance)|false|none|none|
|»» id|integer|false|none|none|
|»» amount|integer|false|none|none|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» err_code|string|true|none|none|
|» message|string|true|none|none|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_user">user</h2>
<!-- backwards compatibility -->
<a id="schemauser"></a>
<a id="schema_user"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": 1,
  "name": "taro"
}

```

user

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer|true|none|none|
|name|string|true|none|none|

<h2 id="tocS_balance_trans">balance_trans</h2>
<!-- backwards compatibility -->
<a id="schemabalance_trans"></a>
<a id="schema_balance_trans"></a>
<a id="tocSbalance_trans"></a>
<a id="tocsbalance_trans"></a>

```json
{
  "id": 12345,
  "user_id": 1,
  "balance_id": 1,
  "transfer": {
    "id": 12345,
    "from_user": 1,
    "from_balance": 1,
    "to_user": 2,
    "to_balance": 2,
    "amount": 1000,
    "processed_at": "2019-08-24T14:15:22Z"
  },
  "amount": 1000,
  "processed_at": "2019-08-24T14:15:22Z"
}

```

balance_trans

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer|true|none|none|
|user_id|integer|true|none|none|
|balance_id|integer|true|none|none|
|transfer|[transer_trans](#schematranser_trans)|false|none|none|
|amount|integer|true|none|none|
|processed_at|string(date-time)|true|none|none|

<h2 id="tocS_transer_trans">transer_trans</h2>
<!-- backwards compatibility -->
<a id="schematranser_trans"></a>
<a id="schema_transer_trans"></a>
<a id="tocStranser_trans"></a>
<a id="tocstranser_trans"></a>

```json
{
  "id": 67890,
  "from_user": 1,
  "from_balance": 1,
  "to_user": 2,
  "to_balance": 2,
  "amount": 1000,
  "processed_at": "2019-08-24T14:15:22Z"
}

```

transer_trans

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer|true|none|none|
|from_user|integer|true|none|none|
|from_balance|integer|true|none|none|
|to_user|integer|true|none|none|
|to_balance|integer|true|none|none|
|amount|integer|false|none|none|
|processed_at|string(date-time)|false|none|none|

<h2 id="tocS_balance">balance</h2>
<!-- backwards compatibility -->
<a id="schemabalance"></a>
<a id="schema_balance"></a>
<a id="tocSbalance"></a>
<a id="tocsbalance"></a>

```json
{
  "id": 1,
  "amount": 5000
}

```

balance

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer|false|none|none|
|amount|integer|false|none|none|

