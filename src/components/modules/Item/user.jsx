import React from "react"
import styles from "./style.css"

const Item = ({ member }) => (
  <li className={styles.userList}>
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