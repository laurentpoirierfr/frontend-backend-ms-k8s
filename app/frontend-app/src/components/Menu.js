import './Menu.css';

function Menu() {


  return (

    <div id="container">
      <ul id="menu">
        <li><a href="#">About Me</a>
          <ul>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas lacinia sem</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
          </ul>
        </li>
        <li><a href="#">Portfolio</a>
          <ul>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas dignissim fermentum luctus</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas lacinia sem</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
          </ul>
        </li>
        <li><a href="#">Clients</a>
          <ul>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas lacinia sem</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
          </ul>
        </li>
        <li><a href="#">Contact Me</a>
          <ul>
            <li><a href="#">Lorem ipsum dolor</a></li>
            <li><a href="#">Maecenas dignissim fermentum luctus</a></li>
            <li><a href="#">Suspendisse fringilla</a></li>
          </ul>
        </li>
        <li><a href="#">Support</a></li>
      </ul>
    </div>
  );
}

export default Menu;
