import React from "react"
import styles from "./style.css"

const Item = ({ member, onSelectHandler }) => (
  <li className={styles.userList} style={{ opacity: member.selected ? 1 : 0.4 }} onClick={onSelectHandler} data-id={member.id}>
    <div className={styles.icon}>
      <img className={styles.iconImage} src={member.profile.image_192} alt="icon image"/>
    </div>
    <div className={styles.names}>
      <p className={styles.realName}>{member.real_name}</p>
      <p className={styles.userId}>@{member.name}</p>
    </div>
  </li>
)

export default Item;