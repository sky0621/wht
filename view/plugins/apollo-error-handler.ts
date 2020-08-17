export default function apolloErrorHandler(
  { graphQLErrors, networkError, operation, forward }: any,
  nuxtContext: any
) {
  console.log('apolloErrorHandler!')

  if (graphQLErrors) {
    console.log(graphQLErrors)
  }
  if (networkError) {
    console.log(networkError)
  }
  if (operation) {
    console.log(operation)
  }
  if (forward) {
    console.log(forward)
  }
  // console.info(nuxtContext)
  nuxtContext.error({ statusCode: 304, message: 'Server error' })
}
