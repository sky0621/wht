import ApolloClient from 'apollo-boost'
import { DefaultApolloClient } from '@vue/apollo-composable'
import { provide } from '@vue/composition-api'

const apolloClient = new ApolloClient({
  connectToDevTools: true,
  uri: 'http://localhost:8080/query',
})
export default function (ctx: any) {
  ctx.app.setup = () => {
    provide(DefaultApolloClient, apolloClient)
  }
}
