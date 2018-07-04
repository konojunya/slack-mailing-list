import React from "react"
import styles from "./style.css"
import client from "../../../api"

export default class ChatArea extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      value: ""
    }
  }

  render() {
    return (
      <div className={styles.wrapper}>
        <textarea cols="30" rows="10" className={styles.textarea} value={this.state.value} onChange={this.onTextChange.bind(this)}></textarea>
        <button className={styles.sendButton} onClick={this.send.bind(this)}>送信</button>
      </div>
    )
  }

  onTextChange(e) {
    this.setState({
      value: e.target.value
    })
  }

  async send() {
    const res = await client.postMessage(this.state.value, this.props.sendTo.map(to => to.id))
    if(res.status === 200) {
      this.setState({
        value: ""
      })
      this.props.clear()
    } else {
      alert("エラーが発生しました。")
    }
  }
}