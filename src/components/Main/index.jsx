import React from "react"
import styles from "./style.css"
import { ApolloProvider, graphql } from 'react-apollo'
import gql from "graphql-tag";
import { ApolloClient } from "apollo-boost";

const client = new ApolloClient({
  uri: `${location.origin}/graphql`
})

const UsersSearchQuery = gql`
  query {
    users
  }
`

class Users extends React.Component {
  render() {
    let users = []
    if(this.props.data) {
      users = this.props.data.users
    }

    return (
      <ul>
        {users.map((user, i) => <li key={index}>hoge</li>)}
      </ul>
    )
  }
}

export default class Main extends React.Component {
  
  state = {
    users: null
  }

  componentWillMount() {
    const UsersWithQuery = graphql(UsersSearchQuery)(Users)
    this.setState({
      users: <UsersWithQuery />
    })
  }

  render() {
    return (
      <ApolloProvider client={client}>
        {this.state.users}
      </ApolloProvider>
    )
  }
}