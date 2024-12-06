# backend/Dockerfile
FROM mysql:8.0

# 文字コードの設定
ENV LANG=C.UTF-8

# タイムゾーンの設定
ENV TZ=Asia/Tokyo

# MySQLの設定ファイルをコピー
COPY ./mysql/my.cnf /etc/mysql/conf.d/my.cnf

# 初期化SQLをコピー
COPY ./mysql/init.sql /docker-entrypoint-initdb.d/

# ポートの公開
EXPOSE 3306