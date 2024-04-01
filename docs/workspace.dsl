workspace {
    name "Сайт Конференции"
    description "Сайт докладов и конференций"

    !identifiers hierarchical

    !adrs decisions

    model {
        user = person "Пользователь сайта конференций"
        conference = softwareSystem "Сайт конференции" {
            description "Сайт для выставления информации о конференциях и докладах на них"

            api_gateway = container "api_gateway" {
                description "Входной API для обращения к основынм сервисам приложения"
            }

            conferences_service = container "conferences service" {
                description "API для получения и управления конференциями и докладами"
            }
            
            users_service = container "users service" {
                description "API для получения и управления пользователями"
            }

            group "Слой хранения данных" {
                user_database = container "User Database" {
                    description "База данных для хранения пользователей"
                    technology "containerized PostreSQL 16.2"
                    tags "database"
                }
                
                cache = container "Cache" {
                    description "Кэш для ускоренного поиска по пользователям"
                    technology "containerized Redis 7.2"
                    tags "database"
                }

                conference_database = container "Conference Database" {
                    description "База для хранения конференция и докладов"
                    technology "MongoDB 7"
                    tags "database"
                }
            }

            user -> api_gateway "API команды на управление и получение данных о конференциях, докладах и пользователях"
            api_gateway -> conferences_service "внутренние API команды на управление и получение данных о конференциях, докладах"
            api_gateway -> users_service "внутренние API команды на управление и получение данных о пользователях"

            conferences_service -> conference_database "Получение/обновление данных о конференциях и докладах" "TCP 27017"
            users_service -> user_database "Получение/обновление данных о пользователях" "TCP 5432"
            users_service -> cache "Получение/обновление данных о пользователях" "TCP 6379"
        }

        user -> conference "Управление и получение данных о конференциях, докладах и пользователях"

        deploymentEnvironment "Production" {
            deploymentNode "Conference server" {
                containerInstance conference.conferences_service
                instances 1
                properties {
                    "cpu" "4"
                    "ram" "4Gb"
                }
            }
            
            deploymentNode "User server" {
                containerInstance conference.users_service
                instances 1
                properties {
                    "cpu" "4"
                    "ram" "4Gb"
                }
            }

            deploymentNode "Databases" {
                deploymentNode "User Database Server" {
                    containerInstance conference.user_database
                    instances 1
                }

                deploymentNode "Cache Server" {
                    containerInstance conference.cache
                    instances 1
                }


                deploymentNode "Conference Database Server" {
                    containerInstance conference.conference_database
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
            user -> conference.api_gateway "Создать нового пользователя (POST /user)"
            conference.api_gateway -> conference.users_service "Создать нового пользователя (POST /user)"
            conference.users_service -> conference.user_database "Сохранить данные о пользователе" 
        }

        dynamic conference "UC02" "Поиск пользователя по логину" {
            autoLayout
            user -> conference.api_gateway "Получить пользоватлея по логину (GET /user/login/{login})"
            conference.api_gateway -> conference.users_service "Получить пользоватлея по логину (GET /user/login/{login})"
            conference.users_service -> conference.user_database "Получить данные о пользователе" 
        }
        
        dynamic conference "UC03" "Поиск пользователя по маске имя и фамилия" {
            autoLayout
            user -> conference.api_gateway "Получить пользователя по ФИ (GET /user/name/{name})"
            conference.api_gateway -> conference.users_service "Получить пользователя по ФИ (GET /user/name/{name})"
            conference.users_service -> conference.cache "Получить данные о пользователе, если они закэшированы" 
            conference.users_service -> conference.user_database "Получить данные о пользователе, если их не закэшированы" 
            conference.users_service -> conference.cache "Записать полученные данные о пользователе" 
        }
         
        dynamic conference "UC04" "Создание доклада" {
            autoLayout
            user -> conference.api_gateway "Добавить доклад (POST /lecture)"
            conference.api_gateway -> conference.conferences_service "Добавить доклад (POST /lecture)"
            conference.conferences_service -> conference.conference_database "Сохранить данные о докладе" 
        }

        dynamic conference "UC05" "Получение списка всех докладов" {
            autoLayout
            user -> conference.api_gateway "Получить доклады (GET /lectures)"
            conference.api_gateway -> conference.conferences_service "Получить доклады (GET /lectures)"
            conference.conferences_service -> conference.conference_database "Получить данные о всех докладах" 
        }
        
        dynamic conference "UC06" "Добавление доклада в конференцию" {
            autoLayout
            user -> conference.api_gateway "Добавить конференцию (POST /conference)"
            conference.api_gateway -> conference.conferences_service "Добавить конференцию (POST /conference)"
            conference.conferences_service -> conference.conference_database "Сохранить данные о конференции" 
        }

        dynamic conference "UC07" "Получение списка докладов в конференции" {
            autoLayout
            user -> conference.api_gateway "Получить доклады конференции (GET /conference/{conference_id}/lectures)"
            conference.api_gateway -> conference.conferences_service "Получить доклады конференции (GET /conference/{conference_id}/lectures)"
            conference.conferences_service -> conference.conference_database "Получить данные о всех докладах конференции" 
        }

        styles {
            element "database" {
                shape cylinder
            }
        }
    }
}