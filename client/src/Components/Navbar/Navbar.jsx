import React, { useContext } from 'react';
import { UserContext } from '../../UserContext';
import './navbar.scss';
function Navbar() {


  // let user = useContext(UserContext)
  // let username
  // if(localStorage.getItem("accessToken") !== "" ){

  //   // user = JSON.parse(user)
  //   username  = user.name
  // }

  const {user, setUser} =  useContext(UserContext)
  let username = user.name
  

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
