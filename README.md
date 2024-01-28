# LanguageSwitch

утилита добавляет в систему хук для перехвата событий нажатия клавишь

поддерживается только Mac

## Установка

###  клонировать репозиторий
```
cd /opt/
git clone https://github.com/voldmir/LanguageSwitch.git
cd LanguageSwitch
```

###  собрать утилиту
```
go build language_switch.go
```

###  копировать конфиг запуска
```
cp ru.my.LaguageSwitch.plist /Library/LaunchAgents/
```

###  Разрешить автозапуск

Системные настройки -> Основные -> Объекты входа -> разрешить в фоновом режиме -> language_switch -> On

Системные настройки -> Конфиденциальность и безопасность -> Универсальный доступ -> language_switch -> On
