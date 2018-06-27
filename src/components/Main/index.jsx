import React from "react"
import styles from "./style.css"
import client from "../../api"

// components
import Header from "../modules/Header"

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
          <li key={member.id} style={{margin: "10px 0"}}>
            <p>name: {member.name}</p>
            <p>real name: {member.real_name}</p>
            <div style={{width: "50px"}}>
              <img src={member.profile.image_192} alt="icon" style={{display: "block", width: "100%"}}/>
            </div>
          </li>
        ))}
        {this.state.channel.list.map(channel => (
          <li key={channel.id} style={{margin: "10px 0"}}>
            <p>{channel.name}</p>
          </li>
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