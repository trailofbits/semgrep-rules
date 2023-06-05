// OK: Lacks 'csrfPrevention: true', but on v4 this option is false by default
//ok: apollo-graphql-v4-csrfPrevention
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
});

// Good: Has 'csrfPrevention: true'
//ok: apollo-graphql-v4-csrfPrevention
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: true,
});

// BAD: Has 'csrfPrevention: false'
//ruleid: apollo-graphql-v4-csrfPrevention
const apollo_server_2 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: false,
});
