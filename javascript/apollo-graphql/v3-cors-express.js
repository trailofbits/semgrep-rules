import { ApolloServer } from 'apollo-server-express';

// ====================
// v3-express-no-cors tests
// ====================
// BAD: Lacks 'cors'
//ruleid: v3-express-no-cors
const apollo_server_1 = new ApolloServer({
    typeDefs,
    resolvers,
});

// GOOD: CORS is defined to deny all
const apollo_server_2 = new ApolloServer({
    typeDefs,
    resolvers,
});
//ok: v3-express-no-cors
apollo_server_2.applyMiddleware({
    app,
    cors: { origin: false },
    path: "/graphql"
});

// ====================
// v3-express-bad-cors tests
// ====================
// GOOD: CORS is defined to deny all
const apollo_server_3 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_3.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: false },
    path: "/graphql"
});

// GOOD: CORS is defined to deny all from var
const apollo_server_4 = new ApolloServer({
    typeDefs,
    resolvers,
});
const apollo_server_4_var = { origin: false }
apollo_server_4.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: apollo_server_4_var,
    path: "/graphql"
});

// GOOD: CORS is defined to deny all from var before server creation
const apollo_server_5_var = { origin: false }
const apollo_server_5 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_5.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: apollo_server_5_var,
    path: "/graphql"
});

// ====================
// BAD: Has a very permissive 'cors'
const apollo_server_6 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_6.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: true },
    path: "/graphql"
});

// BAD: Has a very permissive 'cors' from a variable after
const apollo_server_7 = new ApolloServer({
    typeDefs,
    resolvers,
});
const apollo_server_7_var = { origin: true }
apollo_server_7.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: apollo_server_7_var,
    path: "/graphql"
});

// BAD: Has a very permissive 'cors' from a variable before server creation
const apollo_server_8_var = { origin: true }
const apollo_server_8 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_8.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: apollo_server_8_var,
    path: "/graphql"
});

// ====================
// BAD: Origins defined with the null origin
const apollo_server_8_1 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_8_1.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: 'null' },
    path: "/graphql"
});

// BAD: Origins defined with the null origin in an array
const apollo_server_8_2 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_8_2.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: ['test.com', 'null', 'test2.com'] },
    path: "/graphql"
});

// ====================
// BAD: Origin defined as a regex with an unescaped '.'
const apollo_server_9 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_9.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: /\.example.com$/ },
    path: "/graphql"
});

// BAD: Origin defined as a regex with a missing '$'
const apollo_server_10 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_10.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: /\.example\.com/ },
    path: "/graphql"
});

// BAD: Origins defined as an array with a regex with unescaped '.'
const apollo_server_11 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_11.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: ['test.com', /\.example.com$/, 'test2.com'] },
    path: "/graphql"
});

// BAD: Origins defined as an array with a regex missing a '$'
const apollo_server_12 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_12.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: ['test.com', /\.example\.com/, 'test2.com'] },
    path: "/graphql"
});

// BAD: Origins defined as an array with regexes with an unescaped '.' and missing '$'
const apollo_server_13 = new ApolloServer({
    typeDefs,
    resolvers,
});
apollo_server_13.applyMiddleware({
    app,
    //ruleid: v3-express-bad-cors
    cors: { origin: ['test.com', /\.example\.com$/, /\.example2.com/] },
    path: "/graphql"
});

// ====================
// GOOD: Origin defined as a string
const apollo_server_14 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_14.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: 'https://example.com' },
    path: "/graphql"
});

// GOOD: Origin defined as an array
const apollo_server_15 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_15.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: ['https://example.com'] },
    path: "/graphql"
});

// GOOD: Origins defined as an array of strings
const apollo_server_16 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_16.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: ['https://example.com', 'test.com', 'test2.com'] },
    path: "/graphql"
});

// GOOD: Origin defined as regex
const apollo_server_17 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_17.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: /\.example2\.com$/ },
    path: "/graphql"
});

// GOOD: Origins defined as an array of strings and a regex
const apollo_server_18 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_18.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: { origin: ['https://example.com',  /\.example2\.com$/, "test.com"] },
    path: "/graphql"
});

// GOOD: Origin defined conditionally
let apollo_server_19_var;
if ((process.env.ENABLE_CORS || 'true') === 'true') {
    const origin = process.env.UI_SERVER_ORIGIN || 'http://localhost:9000';
    const methods = 'POST';
    apollo_server_19_var = { origin, methods, credentials: true };
} else {
    apollo_server_19_var = { origin: false };
}
const apollo_server_19 = new ApolloServer({
    typeDefs,
    resolvers,
}); 
apollo_server_19.applyMiddleware({
    app,
    //ok: v3-express-bad-cors
    cors: apollo_server_19_var,
    path: "/graphql"
});