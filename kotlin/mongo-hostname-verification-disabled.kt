package test

import com.mongodb.MongoClientSettings
import com.mongodb.MongoClients

fun main() {
    val mongoClient = MongoClient.create(
        MongoClientSettings.builder()
            .applyConnectionString(ConnectionString("<your connection string>"))
            .applyToSslSettings{ builder ->
                builder.enabled(true)
                // ruleid: mongo-hostname-verification-disabled
                builder.invalidHostNameAllowed(true);
            }
            .build()
    )

    val mongoClient = MongoClient.create(
        MongoClientSettings.builder()
            .applyConnectionString(ConnectionString("<your connection string>"))
            .applyToSslSettings{ builder ->
                // ok: mongo-hostname-verification-disabled
                builder.enabled(true)
            }
            .build()
    )
}
