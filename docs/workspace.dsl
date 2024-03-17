workspace {
    name "Сайт Конференции"
    description "Сайт докладов и конференций"

    !identifiers hierarchical

    !adrs decisions

    # Модель архитектуры
    model {
        # Подумать, возможно должен быть пользователь, который просто может смотреть и администратор,
        # который будет добавлять, но тогда нужна ролевая модель, но с другой стороны она может быть на другом уровне
        user = person "Пользователь сайта конференций"
        conference = softwareSystem "Сайт конференции" {
            description "Сайт для выставления информации о конференциях и докладах на них"

            conference_service = container "conference service" {
                description "API для получения и управления конференциями, докладами и пользователями"
            }

            group "Слой хранения данных" {
                database = container "Database" {
                    description "База данных для хранения конференций докладов и пользователей"
                    technology "containerized PostreSQL 16.2"
                    tags "database"
                }
                
                cache = container "Cache" {
                    description "Кэш для ускоренного поиска по пользователям"
                    technology "containerized Redis 7.2"
                    tags "database"
                }
            }

            user -> conference_service "API команды на управление и получение данных о конференциях, докладах и пользователях"

            conference_service -> database "Получение/обновление данных" "TCP 5432"
            conference_service -> cache "Получение/обновление данных о пользователях" "TCP 6379"
        }

        user -> conference "Управление и получение данных о конференциях, докладах и пользователях"

        deploymentEnvironment "Production" {
            deploymentNode "Application server" {
                containerInstance conference.conference_service
                instances 1
                properties {
                    "cpu" "4"
                    "ram" "4Gb"
                }
            }

            deploymentNode "Databases" {
                deploymentNode "Database Server" {
                    containerInstance conference.database
                    instances 1
                }

                deploymentNode "Cache Server" {
                    containerInstance conference.cache
                    instances 1
                }
            }
        }
    }

    views {
        themes default

        properties { 
            structurizr.tooltips true
        }

        !script groovy {
            workspace.views.createDefaultViews()
            workspace.views.views.findAll { it instanceof com.structurizr.view.ModelView }.each { it.enableAutomaticLayout() }
        }

        dynamic conference "UC01" "Создание нового пользователя" {
            autoLayout
            user -> conference.conference_service "Создать нового пользователя (POST /user)"
            conference.conference_service -> conference.database "Сохранить данные о пользователе" 
        }

        dynamic conference "UC02" "Поиск пользователя по логину" {
            autoLayout
            user -> conference.conference_service "Получить пользоватлея по логину (GET /user/login/{login})"
            conference.conference_service -> conference.database "Получить данные о пользователе" 
        }
        
        dynamic conference "UC03" "Поиск пользователя по маске имя и фамилия" {
            autoLayout
            user -> conference.conference_service "Получить пользователя по ФИ (GET /user/name/{name})"
            conference.conference_service -> conference.cache "Получить данные о пользователе, если они закэшированы" 
            conference.conference_service -> conference.database "Получить данные о пользователе, если их не закэшированы" 
            conference.conference_service -> conference.cache "Записать полученные данные о пользователе" 
        }
         
        dynamic conference "UC04" "Создание доклада" {
            autoLayout
            user -> conference.conference_service "Добавить доклад (POST /lecture)"
            conference.conference_service -> conference.database "Сохранить данные о докладе" 
        }

        dynamic conference "UC05" "Получение списка всех докладов" {
            autoLayout
            user -> conference.conference_service "Получить доклады (GET /lectures)"
            conference.conference_service -> conference.database "Получить данные о всех докладах" 
        }
        
        dynamic conference "UC06" "Добавление доклада в конференцию" {
            autoLayout
            user -> conference.conference_service "Добавить конференцию (POST /conference)"
            conference.conference_service -> conference.database "Сохранить данные о конференции" 
        }

        dynamic conference "UC07" "Получение списка докладов в конференции" {
            autoLayout
            user -> conference.conference_service "Получить доклады конференции (GET /conference/{conference_id}/lectures)"
            conference.conference_service -> conference.database "Получить данные о всех докладах конференции" 
        }

        styles {
            element "database" {
                shape cylinder
            }
        }
    }
}