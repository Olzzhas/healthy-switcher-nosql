import React from 'react';
import './login.scss';

import axios from 'axios'

function Login({token, setToken}) {
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
 
  async function registerUser(event) {
    event.preventDefault();

    axios.post("http://localhost:4000/api/tokens/authentication",{
        email: email,
        password: password,

    },)
    .then(res=>{
        setToken(res.data.authentication_token.token)
      
    }).then(()=>{
      window.location.assign("http://localhost:3000");
    })
  }

  return (
    <div className="registerWrapper">
      <ul className="register-content">
        <li className="content-left">Registration</li>

        <form onSubmit={(e)=>registerUser(e)} action="">
          <li className="content-right">
            
            <input
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              type="email"
              name="email"
              placeholder="Enter your email address..."
            />
            <input
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              type="password"
              name="password"
              placeholder="Enter your password..."
            />
          

            <button type="submit">Sign Up</button>
          </li>
        </form>
      </ul>
    </div>
  );
}

export default Login;
