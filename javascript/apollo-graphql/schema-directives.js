// BAD: Has 'schemaDirectives'
//ruleid: schema-directives
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
    schemaDirectives: {
        rateLimit: rateLimitDirective
    },
});

// Good: Does not have 'schemaDirectives'
//ok: schema-directives
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
});