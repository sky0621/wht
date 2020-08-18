import gql from 'graphql-tag'

export const FindWht = gql`
  query FindWht($condition: WhtConditionInput) {
    findWht(condition: $condition) {
      id
      recordDate
      title
      textContents {
        id
        name
        text
      }
      imageContents {
        id
        name
        path
      }
    }
  }
`
