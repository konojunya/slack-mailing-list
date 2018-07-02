import React from "react"
import styles from "./style.css"

const Item = ({ channel, onSelectHandler }) => (
  <li className={styles.channelList} style={{ opacity: channel.selected ? 1 : 0.4 }} onClick={onSelectHandler} data-id={channel.id}>
    <p className={styles.channelId}>#{channel.name}</p>
  </li>
)

export default Item;