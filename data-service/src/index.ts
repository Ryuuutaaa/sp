import 'dotenv/config'
import { createSchema, createYoga } from 'graphql-yoga'
import { createServer } from 'node:http'
import { typeDefs } from './graphql/typeDefs/index.js'
import { resolvers } from './graphql/resolvers/index.js'

const yoga = createYoga({
  schema: createSchema({ typeDefs, resolvers }),
  graphqlEndpoint: '/graphql',
})

const server = createServer(yoga)
const port = Number(process.env.PORT) || 4000

server.listen(port, () => {
  console.log(`🚀 Data Service (GraphQL) running at http://localhost:${port}/graphql`)
})