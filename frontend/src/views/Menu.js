import React, { useState, useEffect } from 'react';
import 'styles/menu.scss';
import { NavLink } from 'react-router-dom';
import { RiRefund2Fill, RiHistoryLine } from "react-icons/ri";
import { VscOctoface } from "react-icons/vsc";
import { AiOutlineKey } from "react-icons/ai";

export default function Menu() {
  return (
    <div className="menu">
      <div className="menu__list">
        <NavLink className="menu__item" to="/create" activeClassName="menu__item--active"><VscOctoface/> Create</NavLink>
        <NavLink className="menu__item" to="/detail" activeClassName="menu__item--active"><AiOutlineKey/> Detail</NavLink>
        <NavLink className="menu__item" to="/send-transaction" activeClassName="menu__item--active"><RiRefund2Fill/> Send Transaction</NavLink>
        <NavLink className="menu__item" to="/history" activeClassName="menu__item--active"><RiHistoryLine/> Transaction History</NavLink>
      </div>
    </div>
  )
}