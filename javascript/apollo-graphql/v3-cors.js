import { ApolloServer } from 'apollo-server';

// ====================
// v3-no-cors tests
// ====================
// BAD: Lacks 'cors'
//ruleid: v3-no-cors
const apollo_server_no_cors_bad_1 = new ApolloServer({
    typeDefs,
    resolvers,
});

// GOOD: CORS is defined to deny all
//ok: v3-no-cors
const apollo_server_no_cors_good_1 = new ApolloServer({
    typeDefs,
    resolvers,
    cors: { origin: false }
});


// ====================
// v3-bad-cors tests
// ====================
// GOOD: CORS is defined to deny all
const apollo_server_bad_cors_good_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-no-cors
    cors: { origin: false }
});

// GOOD: CORS is defined to deny all
const apollo_server_bad_cors_good_1_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-no-cors
    cors: { origin: false, credentials: true }
});

// GOOD: CORS is defined to deny all from var
const apollo_server_bad_cors_good_2_var = { origin: false }
const apollo_server_bad_cors_good_2 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-no-cors
    cors: apollo_server_bad_cors_good_2_var
});

// ====================
// BAD: Has a very permissive 'cors'
const apollo_server_bad_cors_bad_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: true }
});

// BAD: Has a very permissive 'cors' from a variable (just the origin)
const apollo_server_bad_cors_bad_1_1_var = true;
const apollo_server_bad_cors_bad_1_1 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: apollo_server_bad_cors_bad_1_1_var }
});

// BAD: Has a very permissive 'cors' from a variable
const apollo_server_bad_cors_bad_2_var = { origin: true }
const apollo_server_bad_cors_bad_2 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: apollo_server_bad_cors_bad_2_var
});

// ====================
// BAD: Origin defined as a regex with an unescaped '.'
const apollo_server_bad_cors_bad_3 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: /\.example.com$/ }
});

// BAD: Origin defined as a regex with a missing '$'
const apollo_server_bad_cors_bad_4 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: /\.example\.com/ }
});

// BAD: Origins defined as an array with a regex with unescaped '.'
const apollo_server_5 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: ['test.com', /\.example.com$/, 'test2.com'] }
});

// BAD: Origins defined as an array with a regex missing a '$'
const apollo_server_6 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: ['test.com', /\.example\.com/, 'test2.com'] }
});

// BAD: Origins defined as an array with regexes with an unescaped '.' and missing '$'
const apollo_server_7 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: ['test.com', /\.example\.com$/, /\.example2.com/] }
});

// BAD: Origins defined with the null origin
const apollo_server_8 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: 'null' }
});

// BAD: Origins defined with the null origin in an array
const apollo_server_9 = new ApolloServer({
    typeDefs,
    resolvers,
    //ruleid: v3-bad-cors
    cors: { origin: ['test.com', 'null', 'test2.com'] }
});

// ====================
// GOOD: Origin defined as a string
const apollo_server_ok_4 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: { origin: 'https://example.com' }
});

// GOOD: Origin defined as an array
const apollo_server_ok_2 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: { origin: ['https://example.com'] }
});

// GOOD: Origins defined as an array of strings
const apollo_server_ok_3 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: { origin: ['https://example.com', 'test.com', 'test2.com'] }
});

// GOOD: Origin defined as regex
const apollo_server_ok_5 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: { origin:  /\.example2\.com$/ }
});

// GOOD: Origins defined as an array of strings and a regex
const apollo_server_ok_6 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: { origin: ['https://example.com',  /\.example2\.com$/, "test.com"] }
});

// GOOD: Origins defined as an array from variable
const apollo_server_ok_8_cors_policy = { origin: ["https://example.com", 'test.com', 'test2.com'] }
const apollo_server_ok_8 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: apollo_server_ok_8_cors_policy
});

// GOOD: Origin defined conditionally
let apollo_server_ok_9_cors_policy;
if ((process.env.ENABLE_CORS || 'true') === 'true') {
    const origin = process.env.UI_SERVER_ORIGIN || 'http://localhost:9000';
    const methods = 'POST';
    apollo_server_ok_9_cors_policy = { origin, methods, credentials: true };
} else {
    apollo_server_ok_9_cors_policy = { origin: false };
}
const apollo_server_ok_9 = new ApolloServer({
    typeDefs,
    resolvers,
    //ok: v3-bad-cors
    cors: apollo_server_ok_9_cors_policy
});

