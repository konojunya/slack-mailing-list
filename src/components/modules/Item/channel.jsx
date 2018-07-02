import React from "react"
import styles from "./style.css"

const Item = ({ name }) => (
  <li className={styles.channelList}>
    <p className={styles.channelId}>#{name}</p>
  </li>
)

export default Item;