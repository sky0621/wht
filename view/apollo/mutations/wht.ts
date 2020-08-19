import gql from 'graphql-tag'

export const CreateWht = gql`
  mutation CreateWht($wht: WhtInput!) {
    createWht(wht: $wht) {
      id
    }
  }
`
