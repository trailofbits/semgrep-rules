// OK: Lacks 'csrfPrevention: true', but on v4 this option is true by default
//ok: v4-csrf-prevention
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
});

// Good: Has 'csrfPrevention: true'
//ok: v4-csrf-prevention
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: true,
});

// BAD: Has 'csrfPrevention: false'
//ruleid: v4-csrf-prevention
const apollo_server_2 = new ApolloServer({
    typeDefs,
    resolvers,
    csrfPrevention: false,
});
