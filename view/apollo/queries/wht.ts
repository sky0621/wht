import gql from 'graphql-tag'

export const FindWht = gql`
  query FindWht {
    findWht {
      id
      recordDate
      path
    }
  }
`
