// BAD 1: Lacks 'csrfPrevention: true'
//ruleid: apollo-graphql-v3-csrfPrevention
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
});

// BAD 2: Has 'csrfPrevention: false'
//ruleid: apollo-graphql-v3-csrfPrevention
const apollo_server_2 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: false,
});

// Good: Has 'csrfPrevention: true'
//ok: apollo-graphql-v3-csrfPrevention
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: true,
});