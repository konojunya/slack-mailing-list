import React from "react"
import client from "../../api"

// components
import Header from "../modules/Header"
import ChatArea from "../modules/ChatArea"
import UserItem from "../modules/Item/user"
import ChannelItem from "../modules/Item/channel"

export default class Main extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      member: {
        list: [],
        nextCursor: ""
      },
      channel: {
        list: [],
        nextCursor: ""
      }
    }
  }

  componentWillMount() {
    this.getUsers()
    this.getChannels()
  }

  render() {
    return (
      <div>
        <Header />
        <ul>
        {this.state.member.list.map(member => (
          <UserItem
            key={member.id}
            member={member}
          />
        ))}
        {this.state.channel.list.map(channel => (
          <ChannelItem
            key={channel.id}
            name={channel.name}
          />
        ))}
        </ul>
      </div>
    )
  }

  async getUsers(nextCursor = "") {
    const res = await client.getUsers(nextCursor)
    this.setState({
      member: {
        list: [...this.state.member.list, ...res.data.members],
        nextCursor: res.data.response_metadata.next_cursor
      }
    })
  }

  async getChannels(nextCursor = "") {
    const res = await client.getChannels(nextCursor)
    this.setState({
      channel: {
        list: [...this.state.channel.list, ...res.data.channels],
        nextCursor: res.data.response_metadata.next_cursor
      }
    })
  }
}