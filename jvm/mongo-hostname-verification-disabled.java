package test

import com.mongodb.MongoClientSettings
import com.mongodb.MongoClients

class HelloWorld {
    public static void main(String[] args) {
        MongoClientSettings settings = MongoClientSettings.builder()
            .applyToSslSettings(builder -> {
                builder.enabled(true);
                // ruleid: mongo-hostname-verification-disabled
                builder.invalidHostNameAllowed(true);
            })
            .build();

        MongoClientSettings settings = MongoClientSettings.builder()
            .applyToSslSettings(builder -> {
                // ok: mongo-hostname-verification-disabled
                builder.enabled(true);
            })
            .build();
    }
}
