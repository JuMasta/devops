# docker-compose deploy
1. git clone https://github.com/JuMasta/devops.git
2. cd docker-compose 
3. docker-compose up

В go-service реализован вебервис на голанг (8080 порт вебсервиса, 8081 для метрик). Метрики - запросы к серверу и голанг процессы. docker-compose сделал через  nginx.
nginx роутит трафик на 8080 вебсервис и 3000 - графана. Пароль для графаны можно посмотреть и поменять в .env.

# helm deploy
1. git clone https://github.com/JuMasta/devops.git
2. cd helm/awesome-service/
3. helm install  awesome-service -f values.yaml .

данный хелм чарт деплоит деплоймент и сервис для работы вебсервиса. промитиус оператора, все необходимые crd, промитиус и графану.
также деплоится ServiceMonitor который натравлен на эндопоинт метрик вебсервиса. Также встроены дашборды кастомные (запросы к серверу и голанг процессы). Проверил на eks кластере отработало успешно.

