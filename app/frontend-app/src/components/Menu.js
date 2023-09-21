import './Menu.css';

function Menu() {

  return (

    <div id="container">
      <ul id="menu">
        <li><a href="#">Activités</a>
          <ul>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas lacinia sem</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
          </ul>
        </li>
        <li><a href="#">Documentation</a>
          <ul>
            <li><a href="#">Fonctionnelle</a></li>
            <li><a href="#">Technique</a></li>
          </ul>
        </li>
        <li><a href="#">Login</a></li>
      </ul>
    </div>
  );
}

export default Menu;
