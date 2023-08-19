データベースを
`docker-compose up`で立ち上げる
[DocekrCompose](https://github.com/SEKI-YUTA/Docker_DockerComposeArchive/tree/master/postgreSQL_CocktailDB)

## ENDPOINTS

-   /ingredients
    -   材料の一覧を取得
-   /cocktails
    -   パラメーター ingredients[]に渡された材料で作れるカクテルの一覧を取得
    -   URL の例：/cocktails?ingredients[]=ジンジャーエール&ingredients[]=ジン
