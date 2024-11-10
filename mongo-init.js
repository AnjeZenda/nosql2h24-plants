db = db.getSiblingDB('plants_market');

db.plants.insertMany([
    {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4d"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
        "image": "https://images.unsplash.com/photo-1567306226416-28f0efdc88ce?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDF8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Маленькие (до 20 см)",
        "price": 500,
        "light_condition": "Тенелюбивые",
        "watering_frequency": "Средний полив (1-2 раза в неделю)",
        "temperature_regime": "Холодостойкие (до 15°C)",
        "care_complexity": "Для начинающих",
        "description": "Растение для небольших тенистых пространств, легкое в уходе.",
        "type": "Комнатное растение",
        "species": "Фикус",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d10"),
        "created_at": ISODate("2024-01-31T10:00:00Z"),
        "place": "Москва"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4e"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
        "image": "https://images.unsplash.com/photo-1501004318641-b39e6451bec6?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDJ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Средние (от 20 до 50 см)",
        "price": 1500,
        "light_condition": "Полутеневые",
        "watering_frequency": "Частый полив (ежедневно)",
        "temperature_regime": "Средний режим (15-22°C)",
        "care_complexity": "Требуют среднего ухода",
        "description": "Универсальное растение для офиса и дома.",
        "type": "Суккулент",
        "species": "Алоэ",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d11"),
        "created_at": ISODate("2024-01-02T11:00:00Z"),
        "place": "Санкт-Петербург"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4f"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
        "image": "https://images.unsplash.com/photo-1504198458649-3128b932f49b?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDR8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Большие (более 50 см)",
        "price": 3000,
        "light_condition": "Светолюбивые",
        "watering_frequency": "Редкий полив (раз в 2 недели)",
        "temperature_regime": "Теплолюбивые (более 22°C)",
        "care_complexity": "Для опытных цветоводов",
        "description": "Эффектное растение для больших помещений.",
        "type": "Пальма",
        "species": "Финиковая пальма",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d12"),
        "created_at": ISODate("2024-09-29T12:00:00Z"),
        "place": "Казань"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c50"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7c"),
        "image": "https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDV8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Маленькие (до 20 см)",
        "price": 800,
        "light_condition": "Полутеневые",
        "watering_frequency": "Средний полив (1-2 раза в неделю)",
        "temperature_regime": "Средний режим (15-22°C)",
        "care_complexity": "Для начинающих",
        "description": "Неприхотливое растение, подходит для небольших помещений.",
        "type": "Комнатное растение",
        "species": "Сансевиерия",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d13"),
        "created_at": ISODate("2024-10-04T13:00:00Z"),
        "place": "Новосибирск"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c51"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7c"),
        "image": "https://images.unsplash.com/photo-1529516532344-1a34b1df1f19?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDZ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Средние (от 20 до 50 см)",
        "price": 1200,
        "light_condition": "Тенелюбивые",
        "watering_frequency": "Частый полив (ежедневно)",
        "temperature_regime": "Средний режим (15-22°C)",
        "care_complexity": "Требуют среднего ухода",
        "description": "Растение для офиса и дома, легкое в уходе.",
        "type": "Папоротник",
        "species": "Нефролепис",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d14"),
        "created_at": ISODate("2024-02-11T14:00:00Z"),
        "place": "Екатеринбург"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c52"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c80"),
        "image": "https://images.unsplash.com/photo-1465101162946-4377e57745c3?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDd8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Большие (более 50 см)",
        "price": 2700,
        "light_condition": "Светолюбивые",
        "watering_frequency": "Редкий полив (раз в 2 недели)",
        "temperature_regime": "Теплолюбивые (более 22°C)",
        "care_complexity": "Для опытных цветоводов",
        "description": "Красивое растение для светлых помещений.",
        "type": "Пальма",
        "species": "Кокосовая пальма",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d15"),
        "created_at": ISODate("2024-08-27T15:00:00Z"),
        "place": "Нижний Новгород"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c53"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c80"),
        "image": "https://images.unsplash.com/photo-1517683059670-9cb6f6b43a59?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDh8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Маленькие (до 20 см)",
        "price": 600,
        "light_condition": "Полутеневые",
        "watering_frequency": "Средний полив (1-2 раза в неделю)",
        "temperature_regime": "Средний режим (15-22°C)",
        "care_complexity": "Для начинающих",
        "description": "Легкое в уходе растение для небольших помещений.",
        "type": "Суккулент",
        "species": "Каланхоэ",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d16"),
        "created_at": ISODate("2024-07-07T16:00:00Z"),
        "place": "Воронеж"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c54"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c83"),
        "image": "https://images.unsplash.com/photo-1526744226395-7d63a3c6a5bb?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDl8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Средние (от 20 до 50 см)",
        "price": 1400,
        "light_condition": "Тенелюбивые",
        "watering_frequency": "Частый полив (ежедневно)",
        "temperature_regime": "Холодостойкие (до 15°C)",
        "care_complexity": "Требуют среднего ухода",
        "description": "Эффектное растение, отлично очищает воздух.",
        "type": "Комнатное растение",
        "species": "Диффенбахия",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d17"),
        "created_at": ISODate("2024-03-30T17:00:00Z"),
        "place": "Самара"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c55"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c83"),
        "image": "https://images.unsplash.com/photo-1511546395754-38b13532eabc?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDEwfHxwbGFudHxlbnwwfHx8fDE2OTk0MTY4OTc&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Большие (более 50 см)",
        "price": 3200,
        "light_condition": "Светолюбивые",
        "watering_frequency": "Редкий полив (раз в 2 недели)",
        "temperature_regime": "Теплолюбивые (более 22°C)",
        "care_complexity": "Для опытных цветоводов",
        "description": "Эффектное растение, требует яркого освещения.",
        "type": "Пальма",
        "species": "Бамбуковая пальма",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d18"),
        "created_at": ISODate("2024-03-22T18:00:00Z"),
        "place": "Челябинск"
      },
      {
        "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c56"),
        "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7e"),
        "image": "https://images.unsplash.com/photo-1501004318641-b39e6451bec6?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDJ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
        "size": "Средние (от 20 до 50 см)",
        "price": 1300,
        "light_condition": "Полутеневые",
        "watering_frequency": "Средний полив (1-2 раза в неделю)",
        "temperature_regime": "Средний режим (15-22°C)",
        "care_complexity": "Требуют среднего ухода",
        "description": "Подходит для создания уюта и комфорта в интерьере.",
        "type": "Комнатное растение",
        "species": "Плющ",
        "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d19"),
        "created_at": ISODate("2024-09-11T19:00:00Z"),
        "place": "Ростов-на-Дону"
      }
])

db.care_rules.insertMany([
    {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d10"),
        "species": "Фикус",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7d"),
            "description_addition": "Поддерживайте влажность воздуха, но избегайте сквозняков.",
            "created_at": ISODate("2023-12-12T10:00:00Z")
          },
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7f"),
            "description_addition": "Разместите на светлом месте, но без прямого солнца.",
            "created_at": ISODate("2024-01-05T14:00:00Z")
          }
        ],
        "created_at": ISODate("2023-12-12T10:00:00Z"),
        "updated_at": ISODate("2024-01-05T14:00:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d11"),
        "species": "Алоэ",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7c"),
            "description_addition": "Поливайте только после высыхания почвы.",
            "created_at": ISODate("2024-02-10T08:30:00Z")
          }
        ],
        "created_at": ISODate("2024-02-10T08:30:00Z"),
        "updated_at": ISODate("2024-02-10T08:30:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d12"),
        "species": "Финиковая пальма",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7e"),
            "description_addition": "Не требует частого полива, оптимально раз в месяц.",
            "created_at": ISODate("2024-04-01T09:00:00Z")
          },
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7b"),
            "description_addition": "Подходит для любых условий освещения.",
            "created_at": ISODate("2024-04-15T13:45:00Z")
          }
        ],
        "created_at": ISODate("2024-04-01T09:00:00Z"),
        "updated_at": ISODate("2024-04-15T13:45:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d13"),
        "species": "Сансевиерия",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c80"),
            "description_addition": "Регулярно протирайте листья влажной тряпкой.",
            "created_at": ISODate("2024-05-05T16:20:00Z")
          },
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7f"),
            "description_addition": "Держите в тепле, избегайте температур ниже 18°C.",
            "created_at": ISODate("2024-06-01T18:10:00Z")
          }
        ],
        "created_at": ISODate("2024-05-05T16:20:00Z"),
        "updated_at": ISODate("2024-06-01T18:10:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d14"),
        "species": "Нефролепис",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c82"),
            "description_addition": "Поливайте погружением корней в воду.",
            "created_at": ISODate("2024-07-01T09:30:00Z")
          },
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c83"),
            "description_addition": "Разместите в помещении с хорошей циркуляцией воздуха.",
            "created_at": ISODate("2024-07-20T17:25:00Z")
          }
        ],
        "created_at": ISODate("2024-07-01T09:30:00Z"),
        "updated_at": ISODate("2024-07-20T17:25:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d15"),
        "species": "Кокосовая пальма",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
            "description_addition": "Поливайте после полного высыхания верхнего слоя почвы.",
            "created_at": ISODate("2024-08-01T14:15:00Z")
          }
        ],
        "created_at": ISODate("2024-08-01T14:15:00Z"),
        "updated_at": ISODate("2024-08-01T14:15:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d16"),
        "species": "Каланхоэ",
        "description": [
          {
            "description_part": {
              "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7b"),
              "description_addition": "Требует умеренного освещения, не любит прямой свет.",
              "created_at": ISODate("2024-09-01T11:30:00Z")
            }
          }
        ],
        "created_at": ISODate("2024-09-01T11:30:00Z"),
        "updated_at": ISODate("2024-09-01T11:30:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d17"),
        "species": "Диффенбахия",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c83"),
            "description_addition": "Светолюбивое растение, нуждается в обилии солнечного света.",
            "created_at": ISODate("2024-10-01T08:20:00Z")
          }
        ],
        "created_at": ISODate("2024-10-01T08:20:00Z"),
        "updated_at": ISODate("2024-10-01T08:20:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d18"),
        "species": "Бамбуковая пальма",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c81"),
            "description_addition": "Поливайте осторожно, чтобы не замачивать листья.",
            "created_at": ISODate("2024-11-01T09:40:00Z")
          }
        ],
        "created_at": ISODate("2024-11-01T09:40:00Z"),
        "updated_at": ISODate("2024-11-01T09:40:00Z")
      },
      {
        "_id": ObjectId("5f3e8c1d1a9e3f1b1b2c3d19"),
        "species": "Плющ",
        "description": [
          {
            "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7e"),
            "description_addition": "Поливайте мягкой водой, избегая застоя.",
            "created_at": ISODate("2024-12-01T10:30:00Z")
          }
        ],
        "created_at": ISODate("2024-12-01T10:30:00Z"),
        "updated_at": ISODate("2024-12-01T10:30:00Z")
      }
])

db.trades.insertMany([
    {
        "_id": ObjectId("672f5ee65d51777333752da1"),
        "offerer": {
            "name": "Иванов Алексей Петрович",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4d"),
                "species": "Фикус"
            }
        },
        "accepter": {
            "name": "Сидоров Дмитрий Иванович",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7c"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c50"),
                "species": "Сансевиерия"
            }
        },
        "created_at": ISODate("2024-10-10T10:00:00Z"),
        "updated_at": ISODate("2024-10-11T10:00:00Z"),
        "status": 3
    },
    {
        "_id": ObjectId("672f5ee65d51777333752da0"),
        "offerer": {
            "name": "Иванов Алексей Петрович",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4e"),
                "species": "Алоэ"
            }
        },
        "accepter": {
            "name": "Кузнецов Сергей Андреевич",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7e"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c56"),
                "species": "Плющ"
            }
        },
        "created_at": ISODate("2024-02-22T10:00:00Z"),
        "updated_at": ISODate("2024-02-28T10:00:00Z"),
        "status": 1
    },
    {
        "_id": ObjectId("672f5ee65d51777333752da6"),
        "offerer": {
            "name": "Васильев Михаил Павлович",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c80"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c52"),
                "species": "Каланхоэ"
            }
        },
        "accepter": {
            "name": "Иванов Алексей Петрович",
            "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
            "plant": {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4f"),
                "species": "Пальма"
            }
        },
        "created_at": ISODate("2024-08-11T10:00:00Z"),
        "updated_at": ISODate("2024-08-28T10:00:00Z"),
        "status": 2
    }
])

db.users.insertMany([
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
        "photo": "https://example.com/photo1.jpg",
        "surname": "Иванов",
        "name": "Алексей",
        "father_name": "Петрович",
        "phone_number": "+79210001111",
        "email": "ivanov.alexey@example.com",
        "password": "password123",
        "created_at": ISODate("2023-12-15T09:30:00Z"),
        "updated_at": ISODate("2024-01-02T12:45:00Z"),
        "plants": [
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4d"),
                "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
                "image": "https://images.unsplash.com/photo-1567306226416-28f0efdc88ce?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDF8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Маленькие (до 20 см)",
                "price": 500,
                "light_condition": "Тенелюбивые",
                "watering_frequency": "Средний полив (1-2 раза в неделю)",
                "temperature_regime": "Холодостойкие (до 15°C)",
                "care_complexity": "Для начинающих",
                "description": "Растение для небольших тенистых пространств, легкое в уходе.",
                "type": "Комнатное растение",
                "species": "Фикус",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d10"),
                "created_at": ISODate("2024-01-31T10:00:00Z"),
                "place": "Москва"
            },
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4e"),
                "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
                "image": "https://images.unsplash.com/photo-1501004318641-b39e6451bec6?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDJ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Средние (от 20 до 50 см)",
                "price": 1500,
                "light_condition": "Полутеневые",
                "watering_frequency": "Частый полив (ежедневно)",
                "temperature_regime": "Средний режим (15-22°C)",
                "care_complexity": "Требуют среднего ухода",
                "description": "Универсальное растение для офиса и дома.",
                "type": "Суккулент",
                "species": "Алоэ",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d11"),
                "created_at": ISODate("2024-01-02T11:00:00Z"),
                "place": "Санкт-Петербург"
            },
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c4f"),
                "user_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7a"),
                "image": "https://images.unsplash.com/photo-1504198458649-3128b932f49b?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDR8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Большие (более 50 см)",
                "price": 3000,
                "light_condition": "Светолюбивые",
                "watering_frequency": "Редкий полив (раз в 2 недели)",
                "temperature_regime": "Теплолюбивые (более 22°C)",
                "care_complexity": "Для опытных цветоводов",
                "description": "Эффектное растение для больших помещений.",
                "type": "Пальма",
                "species": "Финиковая пальма",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d12"),
                "created_at": ISODate("2024-09-29T12:00:00Z"),
                "place": "Казань"
            }
        ],
        "trades": [
            ObjectId("672f5ee65d51777333752da1"),
            ObjectId("672f5ee65d51777333752da0"),
            ObjectId("672f5ee65d51777333752da6")
        ],
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7b"),
        "photo": "https://example.com/photo2.jpg",
        "surname": "Петров",
        "name": "Виктор",
        "father_name": "Алексеевич",
        "phone_number": "+79210002222",
        "email": "petrov.viktor@example.com",
        "password": "securePass321",
        "created_at": ISODate("2024-01-10T10:15:00Z"),
        "updated_at": ISODate("2024-01-15T11:20:00Z"),
        "role": 1
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7c"),
        "photo": "https://example.com/photo3.jpg",
        "surname": "Сидоров",
        "name": "Дмитрий",
        "father_name": "Иванович",
        "phone_number": "+79210003333",
        "email": "sidorov.dmitry@example.com",
        "password": "pass1234",
        "created_at": ISODate("2024-02-01T08:00:00Z"),
        "updated_at": ISODate("2024-02-10T15:30:00Z"),
        "plants": [
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c50"),
                "image": "https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDV8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Маленькие (до 20 см)",
                "price": 800,
                "light_condition": "Полутеневые",
                "watering_frequency": "Средний полив (1-2 раза в неделю)",
                "temperature_regime": "Средний режим (15-22°C)",
                "care_complexity": "Для начинающих",
                "description": "Неприхотливое растение, подходит для небольших помещений.",
                "type": "Комнатное растение",
                "species": "Сансевиерия",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d13"),
                "created_at": ISODate("2024-10-04T13:00:00Z"),
                "place": "Новосибирск"
            },
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c51"),
                "image": "https://images.unsplash.com/photo-1529516532344-1a34b1df1f19?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDZ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Средние (от 20 до 50 см)",
                "price": 1200,
                "light_condition": "Тенелюбивые",
                "watering_frequency": "Частый полив (ежедневно)",
                "temperature_regime": "Средний режим (15-22°C)",
                "care_complexity": "Требуют среднего ухода",
                "description": "Растение для офиса и дома, легкое в уходе.",
                "type": "Папоротник",
                "species": "Нефролепис",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d14"),
                "created_at": ISODate("2024-02-11T14:00:00Z"),
                "place": "Екатеринбург"
            }
        ],
        "trades": [
            ObjectId("672f5ee65d51777333752da1")
        ],
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7d"),
        "photo": "https://example.com/photo4.jpg",
        "surname": "Алексеев",
        "name": "Николай",
        "father_name": "Фёдорович",
        "phone_number": "+79210004444",
        "email": "alekseev.nikolay@example.com",
        "password": "pass5678",
        "created_at": ISODate("2024-03-05T14:00:00Z"),
        "updated_at": ISODate("2024-03-15T16:45:00Z"),
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7e"),
        "photo": "https://example.com/photo5.jpg",
        "surname": "Кузнецов",
        "name": "Сергей",
        "father_name": "Андреевич",
        "phone_number": "+79210005555",
        "email": "kuznetsov.sergey@example.com",
        "password": "MyStrongPass!",
        "created_at": ISODate("2024-04-01T07:45:00Z"),
        "updated_at": ISODate("2024-04-10T13:50:00Z"),
        "plants": [
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c56"),
                "image": "https://images.unsplash.com/photo-1501004318641-b39e6451bec6?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDJ8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Средние (от 20 до 50 см)",
                "price": 1300,
                "light_condition": "Полутеневые",
                "watering_frequency": "Средний полив (1-2 раза в неделю)",
                "temperature_regime": "Средний режим (15-22°C)",
                "care_complexity": "Требуют среднего ухода",
                "description": "Подходит для создания уюта и комфорта в интерьере.",
                "type": "Комнатное растение",
                "species": "Плющ",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d19"),
                "created_at": ISODate("2024-09-11T19:00:00Z"),
                "place": "Ростов-на-Дону"
            }
        ],
        "trades": [
            ObjectId("672f5ee65d51777333752da0")
        ],
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c7f"),
        "photo": "https://example.com/photo6.jpg",
        "surname": "Новиков",
        "name": "Андрей",
        "father_name": "Викторович",
        "phone_number": "+79210006666",
        "email": "novikov.andrey@example.com",
        "password": "andrey2024",
        "created_at": ISODate("2024-05-15T09:10:00Z"),
        "updated_at": ISODate("2024-05-20T10:30:00Z"),
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c80"),
        "photo": "https://example.com/photo7.jpg",
        "surname": "Васильев",
        "name": "Михаил",
        "father_name": "Павлович",
        "phone_number": "+79210007777",
        "email": "vasilyev.mikhail@example.com",
        "password": "securePass2024",
        "created_at": ISODate("2024-06-01T11:15:00Z"),
        "updated_at": ISODate("2024-06-10T15:25:00Z"),
        "plants": [
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c52"),
                "image": "https://images.unsplash.com/photo-1465101162946-4377e57745c3?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDd8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Большие (более 50 см)",
                "price": 2700,
                "light_condition": "Светолюбивые",
                "watering_frequency": "Редкий полив (раз в 2 недели)",
                "temperature_regime": "Теплолюбивые (более 22°C)",
                "care_complexity": "Для опытных цветоводов",
                "description": "Красивое растение для светлых помещений.",
                "type": "Пальма",
                "species": "Кокосовая пальма",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d15"),
                "created_at": ISODate("2024-08-27T15:00:00Z"),
                "place": "Нижний Новгород"
            },
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c53"),
                "image": "https://images.unsplash.com/photo-1517683059670-9cb6f6b43a59?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDh8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Маленькие (до 20 см)",
                "price": 600,
                "light_condition": "Полутеневые",
                "watering_frequency": "Средний полив (1-2 раза в неделю)",
                "temperature_regime": "Средний режим (15-22°C)",
                "care_complexity": "Для начинающих",
                "description": "Легкое в уходе растение для небольших помещений.",
                "type": "Суккулент",
                "species": "Каланхоэ",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d16"),
                "created_at": ISODate("2024-07-07T16:00:00Z"),
                "place": "Воронеж"
            }
        ],
        "trades": [
            ObjectId("672f5ee65d51777333752da6")
        ],
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c81"),
        "photo": "https://example.com/photo8.jpg",
        "surname": "Макаров",
        "name": "Юрий",
        "father_name": "Сергеевич",
        "phone_number": "+79210008888",
        "email": "makarov.yuriy@example.com",
        "password": "uniquePass098",
        "created_at": ISODate("2024-07-10T08:20:00Z"),
        "updated_at": ISODate("2024-07-15T12:35:00Z"),
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c82"),
        "photo": "https://example.com/photo9.jpg",
        "surname": "Захаров",
        "name": "Игорь",
        "father_name": "Александрович",
        "phone_number": "+79210009999",
        "email": "zakharov.igor@example.com",
        "password": "mypassword_2024",
        "created_at": ISODate("2024-08-15T14:40:00Z"),
        "updated_at": ISODate("2024-08-25T16:55:00Z"),
        "role": 2
    },
    {
        "_id": ObjectId("5f2d8c1d1d8e2f1a1a2b3c83"),
        "photo": "https://example.com/photo10.jpg",
        "surname": "Морозов",
        "name": "Олег",
        "father_name": "Геннадьевич",
        "phone_number": "+79210001010",
        "email": "morozov.oleg@example.com",
        "password": "oleg_secure123",
        "created_at": ISODate("2024-09-05T18:05:00Z"),
        "updated_at": ISODate("2024-09-15T20:30:00Z"),
        "plants": [
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c54"),
                "image": "https://images.unsplash.com/photo-1526744226395-7d63a3c6a5bb?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDl8fHBsYW50fGVufDB8fHx8MTY5OTQxNjg5Nw&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Средние (от 20 до 50 см)",
                "price": 1400,
                "light_condition": "Тенелюбивые",
                "watering_frequency": "Частый полив (ежедневно)",
                "temperature_regime": "Холодостойкие (до 15°C)",
                "care_complexity": "Требуют среднего ухода",
                "description": "Эффектное растение, отлично очищает воздух.",
                "type": "Комнатное растение",
                "species": "Диффенбахия",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d17"),
                "created_at": ISODate("2024-03-30T17:00:00Z"),
                "place": "Самара"
            },
            {
                "_id": ObjectId("5f1d7c1d1d8e2f1a1a2b3c55"),
                "image": "https://images.unsplash.com/photo-1511546395754-38b13532eabc?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDEwfHxwbGFudHxlbnwwfHx8fDE2OTk0MTY4OTc&ixlib=rb-1.2.1&q=80&w=400",
                "size": "Большие (более 50 см)",
                "price": 3200,
                "light_condition": "Светолюбивые",
                "watering_frequency": "Редкий полив (раз в 2 недели)",
                "temperature_regime": "Теплолюбивые (более 22°C)",
                "care_complexity": "Для опытных цветоводов",
                "description": "Эффектное растение, требует яркого освещения.",
                "type": "Пальма",
                "species": "Бамбуковая пальма",
                "care_rules": ObjectId("5f3e8c1d1a9e3f1b1b2c3d18"),
                "created_at": ISODate("2024-03-22T18:00:00Z"),
                "place": "Челябинск"
            }
        ],
        "role": 2
    }
])