import React from "react"
import styles from "./style.css"
import client from "../../api"

// components
import Header from "../modules/Header"

export default class Main extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      members: []
    }
  }

  componentWillMount() {
    this.getUsers()
  }

  render() {
    return (
      <div>
        <Header />
        <ul>
        {this.state.members.map(member => (
          <li key={member.id} style={{margin: "10px 0"}}>
            <p>name: {member.name}</p>
            <p>real name: {member.real_name}</p>
            <div style={{width: "50px"}}>
              <img src={member.profile.image_192} alt="icon" style={{display: "block", width: "100%"}}/>
            </div>
          </li>
        ))}
        </ul>
      </div>
    )
  }

  async getUsers() {
    const res = await client.getUsers()
    this.setState({
      members: res.data.members
    })
  }
}