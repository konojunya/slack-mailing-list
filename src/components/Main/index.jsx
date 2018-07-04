import React from "react"
import client from "../../api"
import styles from "./style.css"

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
        <div className={styles.container}>
          <ul className={styles.memberList}>
          {this.state.member.list.map(member => (
            <UserItem
              key={member.id}
              member={member}
              onSelectHandler={this.onSelectUser.bind(this)}
            />
          ))}
          </ul>
          <div className={styles.channelsAndTextArea}>
            <ul className={styles.channelList}>
              {this.state.channel.list.map(channel => (
                <ChannelItem
                  key={channel.id}
                  channel={channel}
                  onSelectHandler={this.onSelectChannel.bind(this)}
                />
              ))}
            </ul>
            <ChatArea sendTo={this.getSend()} clear={this.clearSelect.bind(this)} />
          </div>
        </div>
      </div>
    )
  }

  getSend() {
    return [
      ...this.state.member.list.filter(member => member.selected),
      ...this.state.channel.list.filter(channel => channel.selected),
    ]
  }

  clearSelect() {
    this.setState({
      member: { list: this.state.member.list.map(member => Object.assign({}, member, { selected: false })) },
      channel: { list: this.state.channel.list.map(channel => Object.assign({}, channel, { selected: false })) },
    })
  }

  onSelectUser(e) {
    const id = e.currentTarget.dataset.id
    this.setState({
      member: {
        list: this.state.member.list.map(member => (
          Object.assign({}, member, {
            selected: member.id === id ? !member.selected : member.selected
          })
        ))
      }
    })
  }
  onSelectChannel(e) {
    const id = e.currentTarget.dataset.id
    this.setState({
      channel: {
        list: this.state.channel.list.map(channel => (
          Object.assign({}, channel, {
            selected: channel.id === id ? !channel.selected : channel.selected
          })
        ))
      }
    })
  }

  async getUsers(nextCursor = "") {
    const res = await client.getUsers(nextCursor)
    this.setState({
      member: {
        list: [
          ...this.state.member.list,
          ...res.data.members.map(member => (
            Object.assign({}, member, {
              selected: false
            })
          ))
        ],
        nextCursor: res.data.response_metadata.next_cursor
      }
    })
  }

  async getChannels(nextCursor = "") {
    const res = await client.getChannels(nextCursor)
    this.setState({
      channel: {
        list: [
          ...this.state.channel.list,
          ...res.data.channels.map(channel => (
            Object.assign({}, channel, {
              selected: false
            })
          ))
        ],
        nextCursor: res.data.response_metadata.next_cursor
      }
    })
  }
}