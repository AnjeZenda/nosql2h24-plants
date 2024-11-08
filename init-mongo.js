db = db.getSiblingDB('plants-market');

db.users.insertMany(
    [
        {
            "_id": "672d0dbdfc13aeaaf467464e",
            "photo": "https://robohash.org/quamiureimpedit.png?size=50x50&set=set1",
            "name": "Аркадий",
            "surname": "Висов",
            "father_name": "Борисович",
            "phone_number": "89399890678",
            "email": "111@mail.ru",
            "password": "12345678",
            "created_at": "1731078721866",
            "updated_at": "1731078721866",
            "plants": [
                {
                    "_id":  "172d0cf2fc13ae0d466745fe",
                    "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "size": "medium",
                    "price": 200.00,
                    "light_condition": "shadow",
                    "watering_frequency": "medium",
                    "temperature_regime": "medium",
                    "care_complexity": "begginer",
                    "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
                    "type": "aloe",
                    "species": "aloe",
                    "care_rules": "672d0ad0fc13ae0c12674602",
                    "created_at": "1731078741866",
                    "place": "Moscow"
                },
                {
                    "_id":  "272d0cf2fc13ae0d466745fe",
                    "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "size": "medium",
                    "price": 200.00,
                    "light_condition": "light",
                    "watering_frequency": "medium",
                    "temperature_regime": "medium",
                    "care_complexity": "begginer",
                    "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
                    "type": "cactus",
                    "species": "cactus",
                    "care_rules": "672d0ad0fc13ae0c126745ff",
                    "created_at": "1731078741866",
                    "place": "Moscow"
                },
                {
                    "_id":  "372d0cf2fc13ae0d466745fe",
                    "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "size": "medium",
                    "price": 200.00,
                    "light_condition": "shadow",
                    "watering_frequency": "medium",
                    "temperature_regime": "medium",
                    "care_complexity": "begginer",
                    "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
                    "type": "yellow rose",
                    "species": "rose",
                    "care_rules": "672d0cf2fc13ae0d46674600",
                    "created_at": "1731078741866",
                    "place": "Moscow"
                }
            ],
            "trades": [
                {
                    "id": "672d114cfc13ae0dea6745fe",
                },
                {
                    "id": "672d114cfc1ffe0dea6745fe",
                }
            ],
            "role": true
        },
        {
            "_id": "672d22bdfc13ae0ef467464e",
            "photo": "https://robohash.org/quamiureimpedit.png?size=50x50&set=set1",
            "name": "Алина",
            "surname": "Кельн",
            "father_name": "Владимировна",
            "phone_number": "89399890911",
            "email": "fsdgdf@mail.ru",
            "password": "87654321",
            "created_at": "1731078721866",
            "updated_at": "1731078721866",
            "plants": [
                {
                    "_id":  "472d0cf2fc13ae0d466745fe",
                    "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
                    "user_id": "672d22bdfc13ae0ef467464e",
                    "size": "medium",
                    "price": 500.00,
                    "light_condition": "semishadow",
                    "watering_frequency": "medium",
                    "temperature_regime": "medium",
                    "care_complexity": "begginer",
                    "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
                    "type": "white rose",
                    "species": "rose",
                    "care_rules": "672d0cf2fc13ae0d46674600",
                    "created_at": "1731078741866",
                    "place": "Moscow"
                }
            ],
            "trades": [
                {
                    "id": "672d114cfc13ae0dea6745fe",
                }
            ],
            "role": false
        },
        {
            "_id": "672d0dbdfc13ae0ef467411e",
            "photo": "https://robohash.org/quamiureimpedit.png?size=50x50&set=set1",
            "name": "Елена",
            "surname": "Тарнова",
            "father_name": "Витальевна",
            "phone_number": "89399230911",
            "email": "qweqwe@mail.ru",
            "password": "gA5%!*6n8",
            "created_at": "1731078721866",
            "updated_at": "1731078721866",
            "plants": [
                {
                    "_id":  "572d0cf2fc13ae0d466745fe",
                    "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
                    "user_id": "672d0dbdfc13ae0ef467411e",
                    "size": "small",
                    "price": 300.00,
                    "light_condition": "shadow",
                    "watering_frequency": "medium",
                    "temperature_regime": "medium",
                    "care_complexity": "professional",
                    "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
                    "type": "black rose",
                    "species": "rose",
                    "care_rules": "672d0cf2fc13ae0d46674600",
                    "created_at": "1731078741866",
                    "place": "Moscow"
                }
            ],
            "trades": [
                {
                    "id": "672d114cfc1ffe0dea6745fe",
                }
            ],
            "role": false
        }
    ]
)

db.trades.insertMany(
    [
        {
            "_id": "672d114cfc13ae0dea6745fe",
            "offerer": {
                "id": "672d0dbdfc13aeaaf467464e",
                "name": "Висов Аркадий Борисович",
                "plant": {
                    "id": "172d0cf2fc13ae0d466745fe",
                    "name": "aloe"
                }
            },
            "acceptor": {
                "id": "672d22bdfc13ae0ef467464e",
                "name": "Кельн Алина Владимировна",
                "plant": {
                    "id": "472d0cf2fc13ae0d466745fe",
                    "name": "white rose"
                }
            },
            "created_at": "1731079586315",
            "updated_at": "1731079586315",
            "status": 1
        },
        {
            "_id": "672d114cfc1ffe0dea6745fe",
            "offerer": {
                "id": "672d0dbdfc13aeaaf467464e",
                "name": "Висов Аркадий Борисович",
                "plant": {
                    "id": "272d0cf2fc13ae0d466745fe",
                    "name": "cactus"
                }
            },
            "acceptor": {
                "id": "672d0dbdfc13ae0ef467411e",
                "name": "Тарнова Елена Витальевна",
                "plant": {
                    "id": "572d0cf2fc13ae0d466745fe",
                    "name": "black rose"
                }
            },
            "created_at": "1731079586315",
            "updated_at": "1731079586315",
            "status": 1
        }
    ]
)

db.plants.insertMany(
    [
        {
            "_id": "172d0cf2fc13ae0d466745fe",
            "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
            "user_id": "672d0dbdfc13aeaaf467464e",
            "size": "medium",
            "price": 200.00,
            "light_condition": "shadow",
            "watering_frequency": "medium",
            "temperature_regime": "medium",
            "care_complexity": "begginer",
            "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
            "type": "aloe",
            "species": "aloe",
            "care_rules": "672d0ad0fc13ae0c12674602",
            "created_at": "1731078741866",
            "place": "Moscow"
        },
        {
            "_id": "272d0cf2fc13ae0d466745fe",
            "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
            "user_id": "672d0dbdfc13aeaaf467464e",
            "size": "medium",
            "price": 200.00,
            "light_condition": "light",
            "watering_frequency": "medium",
            "temperature_regime": "medium",
            "care_complexity": "begginer",
            "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
            "type": "cactus",
            "species": "cactus",
            "care_rules": "672d0ad0fc13ae0c126745ff",
            "created_at": "1731078741866",
            "place": "Moscow"
        },
        {
            "_id": "372d0cf2fc13ae0d466745fe",
            "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
            "user_id": "672d0dbdfc13aeaaf467464e",
            "size": "medium",
            "price": 200.00,
            "light_condition": "shadow",
            "watering_frequency": "medium",
            "temperature_regime": "medium",
            "care_complexity": "begginer",
            "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
            "type": "yellow rose",
            "species": "rose",
            "care_rules": "672d0cf2fc13ae0d46674600",
            "created_at": "1731078741866",
            "place": "Moscow"
        },
        {
            "_id": "472d0cf2fc13ae0d466745fe",
            "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
            "user_id": "672d22bdfc13ae0ef467464e",
            "size": "medium",
            "price": 500.00,
            "light_condition": "semishadow",
            "watering_frequency": "medium",
            "temperature_regime": "medium",
            "care_complexity": "begginer",
            "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
            "type": "white rose",
            "species": "rose",
            "care_rules": "672d0cf2fc13ae0d46674600",
            "created_at": "1731078741866",
            "place": "Moscow"
        },
        {
            "_id": "572d0cf2fc13ae0d466745fe",
            "image": "https://robohash.org/consequaturquodoloribus.png?size=50x50&set=set1",
            "user_id": "672d0dbdfc13ae0ef467411e",
            "size": "small",
            "price": 300.00,
            "light_condition": "shadow",
            "watering_frequency": "medium",
            "temperature_regime": "medium",
            "care_complexity": "professional",
            "description": "justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus",
            "type": "black rose",
            "species": "rose",
            "care_rules": "672d0cf2fc13ae0d46674600",
            "created_at": "1731078741866",
            "place": "Moscow"
        }
    ]
)

db.care_rules.insertMany(
    [
        {
            "_id": "672d0cf2fc13ae0d46674600",
            "species": "rose",
            "description": [
                {
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "description_addition": "phasellus in felis donec semper sapien a libero nam dui proin leo odio porttitor id consequat in consequat ut nulla sed accumsan felis ut at dolor quis odio consequat varius integer",
                    "created_at": "1731079197234"
                },
                {
                    "user_id": "672d22bdfc13ae0ef467464e",
                    "description_addition": "non mattis pulvinar nulla pede nulla sed vel enim sit amet nunc viverra dapibus nulla suscipit ligula in lacus curabitur at ipsum ac tellus semper interdum mauris",
                    "created_at": "1731079197234"
                }
            ],
            "created_at": "1731079191234",
            "updated_at": "1731079191234"
        },
        {
            "_id": "672d0ad0fc13ae0c126745ff",
            "species": "cactus",
            "description": [
                {
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "description_addition": "eget semper rutrum nulla nunc purus phasellus in felis donec semper sapien a libero nam dui proin leo odio porttitor id consequat in consequat ut nulla sed accumsan felis ut at dolor quis odio consequat varius integer",
                    "created_at": "1731079197234"
                },
                {
                    "user_id": "672d22bdfc13ae0ef467464e",
                    "description_addition": "non mattis pulvinar nulla pede ullamcorper augue a suscipit nulla elit ac nulla sed vel enim sit amet nunc viverra dapibus nulla suscipit ligula in lacus curabitur at ipsum ac tellus semper interdum mauris",
                    "created_at": "1731079197234"
                }
            ],
            "created_at": "1731079191234",
            "updated_at": "1731079191234"
        },
        {
            "_id": "672d0ad0fc13ae0c12674602",
            "species": "aloe",
            "description": [
                {
                    "user_id": "672d0dbdfc13aeaaf467464e",
                    "description_addition": "apien a libero nam dui proin leo odio porttitor id consequat in consequat ut nulla sed accumsan felis ut at dolor quis odio consequat varius integer",
                    "created_at": "1731079197234"
                },
                {
                    "user_id": "672d22bdfc13ae0ef467464e",
                    "description_addition": "enim sit amet nunc viverra dapibus nulla suscipit ligula in lacus curabitur at ipsum ac tellus semper interdum mauris",
                    "created_at": "1731079197234"
                }
            ],
            "created_at": "1731079191234",
            "updated_at": "1731079191234"
        }
    ])