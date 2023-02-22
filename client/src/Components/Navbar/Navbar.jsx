import React, { useContext } from 'react';
import { UserContext } from '../../UserContext';
import './navbar.scss';
function Navbar() {
  const user = useContext(UserContext)
  let username = "" 
  if(user !== ""){
    
    // console.log("user is equal to " + user);
    username = user.name
    // console.log("username is equal to " + username);
  }

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
            {localStorage.getItem("accessToken") === "" ? 
            <>
              <li>
              <a href="/login">Login</a>
            </li>
            <li>
              <a href="/register">Sign Up</a>
            </li>
            </>
            :
            <>
            <li>
              <a href="/#">
                  {username}
              </a>
            </li>
              <li>
                <a onClick={()=>{
                  localStorage.setItem("accessToken", "")
                  localStorage.setItem("currentUser", {auth: null})
                  }} href="/">Log Out
                </a>
              </li>
            </>
            }
            
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
