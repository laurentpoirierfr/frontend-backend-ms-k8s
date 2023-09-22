import React, { useEffect, useState } from "react";
import './Menu.css';
import axios from 'axios';

function Menu() {
  const [menu, setMenu] = useState([]);

  useEffect(() => {
    axios.get("/api/menus").then((data) => {
      setMenu(data?.data);
    });
  }, []);

  const TocItem = ({ toc }) => {
    const subItem = (toc.items || []).map(item => (
      <TocItem toc={item} />
    ))

    return (
      <li key={toc.title}>
        <a href={toc.url}>{toc.title}</a>
        {toc.items && toc.items.length &&
          (
            <ul>
              {subItem}
            </ul>
          )
        }
      </li>
    )
  }

  const TableOfContents = ({ toc }) => (
    <ul id="menu">
      {(toc || []).map((item, index) => (
        <TocItem key={index} toc={item} />
      ))}
    </ul>
  )

  return (
    <div id="container">
      <TableOfContents toc={menu} />
    </div>
  );
}

export default Menu;
