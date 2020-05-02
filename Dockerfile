# Test stage

# AS キーワードで、multi-stage build 環境を構築する
# 同一 Dockerfile 内で複数の FROM を定義できるようになる
FROM alpine AS test
LABEL application=money-boy

# バッファを作らない設定
ENV PYTHONUNBUFFERED 1

# Install basic utilities
# apk -> Alpine Linux の package manager
# キャッシュさせないように --no-cache を指定
# Git は Docker にバージョンタグをつけるために付与
RUN apk add --no-cache bash git

# Install build dependencies
# Python C の拡張モジュールを追加
# postgresql-dev は、PostgesSQL client を構築する
# Pillow needs zlib-dev libjpeg package ibjpeg-turbo
RUN apk add --no-cache gcc python3-dev libffi-dev musl-dev linux-headers postgresql-dev zlib-dev jpeg-dev libjpeg-turbo

# 立ち上げ時に、django-markdownx の markdownify が、/usr/lib/libjpeg.so.8 を探しに行くので、以下のリンクを貼る
# ImportError: Error loading shared library libjpeg.so.8: No such file or directory (needed by /usr/lib/python3.6/site-packages/PIL/_imaging.cpython-36m-x86_64-linux-gnu.so)
RUN ln -s /usr/lib/python3.6/site-packages/PIL/_imaging.cpython-36m-x86_64-linux-gnu.so

RUN pip3 install --upgrade pip

# 事前コンパイルおよび事前フォーマットを行うため wheel を導入
RUN pip3 install wheel

# Copy requirements
COPY /src/requirements* /build/
WORKDIR /build

# Build and install requirements
# --no-cahce-dir は、イメージの肥大化を防ぐ
# --no-input でユーザー確認を無効にする
RUN pip3 wheel -r requirements_test.txt --no-cache-dir --no-input

# コンテナに事前ビルドした wheel を導入する
# --no-index は、インターネットからパッケージをダウンロードしないようにする(wheelを使うため)
# -f は、インストール先を /build に作成した wheel に指定
RUN pip3 install -r requirements_test.txt -f /build --no-index --no-cache-dir

# Copy source code
# ソースコードは、アプリケーションの依存関係を先に構築してからインストールする
# 最適化されたキャッシュ環境を構築しているため、イメージが作成されるたびにアプリケーションの依存関係を構築する必要はない
COPY /src /app
WORKDIR /app

# Test entrypoint
# CMD ["python3", "manage.py", "test", "--noinput", "--settings=money-boy.settings_test"]

# Relese stage

FROM alpine
LABEL application=money-boy

# Install operating system dependencies
# Test stage で、プリコンパイルした環境を再利用するので、xxxx-dev 環境は不要
RUN apk add --no-cache python3 postgresql-client bash libjpeg-turbo

# バッファを作らない設定
ENV PYTHONUNBUFFERED 1

# Create app user
# Group ID 1000 のグループを作成
# User ID 1000 のユーザーを作成
# User を Group に追加
# User ログイン時のホームディレクトリを app に指定
# http://raining.bear-life.com/linux/linux%E3%81%AEadduser%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A7%E6%96%B0%E8%A6%8F%E3%83%A6%E3%83%BC%E3%82%B6%E3%82%92%E4%BD%9C%E6%88%90%E3%81%99%E3%82%8B
RUN addgroup -g 1000 app && \
    adduser -u 1000 -G app -D app

# Copy and install application source and pre-built dependencies
# 環境構築用のフォルダを test 環境からコピー
COPY --from=test --chown=app:app /build /build
# テスト環境のソースコードをコピー
COPY --from=test --chown=app:app /app /app
RUN pip3 install -r /build/requirements.txt -f /build --no-index --no-cache-dir
# 環境構築用のフォルダを削除
RUN rm -rf /build

# Create public volume
RUN mkdir /public
RUN chown app:app /public
VOLUME /public

# Set working directory and application user
WORKDIR /app
USER app