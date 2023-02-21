import './navbar.scss';
function Navbar() {
  return (
    <div>
      <div className="navbar">
        <a href="/">
          <div className="logo">
            <img src="/img/Logo.png" alt="logo" />
            <img className="hs-text" src="/img/hs-title.png" alt="title" />
          </div>
        </a>
        <div>
          <ul className="navigation">
            <li>
              <a href="/#">Menu</a>
            </li>
            <li>
              <a href="/#">Recipes</a>
            </li>
            <li>
              <a href="/#">Chefs</a>
            </li>
            <li>
              <a href="/register">Sign Up</a>
            </li>
          </ul>
        </div>
      </div>
      <div>
        <img className="line" src="/img/line.png" alt="line" />
      </div>
    </div>
  );
}

export default Navbar;
