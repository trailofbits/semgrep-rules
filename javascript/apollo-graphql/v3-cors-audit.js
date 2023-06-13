import 'apollo-server';

// GOOD: CORS is defined to deny all
const apollo_graphql_potentially_bad_cors_good_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-potentially-bad-cors
    cors: { origin: false }
});

// GOOD: CORS is defined to deny all with []
const apollo_graphql_potentially_bad_cors_good_2 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-potentially-bad-cors
    cors: { origin: [] }
});

// GOOD: CORS is defined to deny all with [] from var
const apollo_graphql_potentially_bad_cors_good_3_var = { origin: [] }
const apollo_graphql_potentially_bad_cors_good_3 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-potentially-bad-cors
    cors: apollo_graphql_potentially_bad_cors_good_3_var
});

// ====================
// BAD: Has a very permissive 'cors'
const apollo_server_bad_cors_bad_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { origin: true }
});

// BAD: Has a very permissive 'cors' from a variable
const apollo_server_bad_cors_bad_2_var = { origin: true }
const apollo_server_bad_cors_bad_2 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: apollo_server_bad_cors_bad_2_var
});

// ====================
// BAD: string
const apollo_graphql_potentially_bad_cors_bad_3 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { origin: "attacker.com"}
});

// BAD: regex
const apollo_graphql_potentially_bad_cors_bad_4 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { origin: /\.attacker\.com$/ }
});

// BAD: array of strings
const apollo_graphql_potentially_bad_cors_bad_5 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { origin: ["attacker.com", "attacker2.com", "attacker3.com"]}
});


// BAD: array of strings and regex
const apollo_graphql_potentially_bad_cors_bad_6 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { origin: ["attacker.com", /attacker2\.com/, "attacker3.com"]}
});

// BAD: function
const apollo_graphql_potentially_bad_cors_bad_7 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: { 
        origin: function (origin, callback) {
            callback(null, true)
      }
    }
});

// BAD: function from var
const apollo_graphql_potentially_bad_cors_bad_8_var = { 
    origin: function (origin, callback) {
        callback(null, true)
  }
}
const apollo_graphql_potentially_bad_cors_bad_8 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-potentially-bad-cors
    cors: apollo_graphql_potentially_bad_cors_bad_8_var
});

