package main

const hello_page = `
<html
  ><head>
    <title>Refracto | Say hello to your {{index . "name"}} 😎</title>
    <style type="text/css">
      body {
        background-color: #f2f7ff;
        color: #4f4f4f;
        font-family: Arial, sans-serif;
        margin: 0;
        padding: 36px;
        line-height: 1;
        font-size: 14px;
      }

      .section {
        margin-bottom: 36px;
      }
      .section.welcome {
        color: #4f4f4f;
      }
      .section.welcome h1 {
        font-size: 26px;
        background-color: #ffffff;
        padding: 18px 22px 15px 22px;
        margin: 0;
        overflow: hidden;
        border-radius: 20px 20px 0 0;
      }
      .section.welcome h1 strong {
        display: inline-block;
        float: left;
      }
      .section.welcome h1 small {
        display: inline-block;
        float: right;
        text-align: right;
        font-size: 18px;
        padding-top: 4px;
        color: #4f4f4f;
      }
      .section.welcome .article {
        border: 4px solid #ffffff;
        padding: 18px 18px 18px 18px;
        border-radius: 0 0 20px 20px;
      }
      .section.welcome .article h3 {
        font-size: 20px;
        margin: 0 0 18px 0;
      }
      .section.welcome .article a {
        color: #477bdf;
      }
      .section.welcome .article a:visited {
        color: #477bdf;
      }
      .section.welcome .article p {
        font-size: 20px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="section welcome">
        <h1><strong>Refracto | Say hello to your app 😎</strong></h1>
        <div class="article">
          <p>{{index . "name"}} is ready!</p>
        </div>
      </div>
    </div>
  </body></html>

`
