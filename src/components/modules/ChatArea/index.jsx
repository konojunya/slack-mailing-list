import React from "react"
import styles from "./style.css"

export default () => (
  <div className="wrapper">
    <textarea cols="30" rows="10" className={styles.textarea}></textarea>
    <button className={styles.sendButton}>送信</button>
  </div>
)