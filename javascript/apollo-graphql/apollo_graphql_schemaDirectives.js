// BAD: Has 'schemaDirectives'
//ruleid: apollo-graphql-schemaDirectives
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
    schemaDirectives: {
        rateLimit: rateLimitDirective
    },
});

// Good: Does not have 'schemaDirectives'
//ok: apollo-graphql-schemaDirectives
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
});