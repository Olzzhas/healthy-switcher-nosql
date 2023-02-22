import React from 'react';
import './register.scss';

import axios from 'axios'

function Register() {
  const [firstName, setFirstName] = React.useState('');
  const [lastName, setLastName] = React.useState('');
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [repassword, setRepassword] = React.useState('');

  async function registerUser(event) {
    event.preventDefault();

    axios.post("http://localhost:5000/api/registration",{
        first_name: firstName,
        last_name: lastName,
        email: email,
        password: password,
        repassword: repassword,
    },)
    .then(res=>{
        console.log(res.data);
    })
  }

  return (
    <div className="registerWrapper">
      <ul className="register-content">
        <li className="content-left">Registration</li>

        <form onSubmit={(e)=>registerUser(e)} action="">
          <li className="content-right">
            <div className="full-name">
              <input
                value={firstName}
                onChange={(e) => setFirstName(e.target.value)}
                type="text"
                name="first_name"
                placeholder="Enter your first name..."
              />
              <input
                value={lastName}
                onChange={(e) => setLastName(e.target.value)}
                type="text"
                name="last_name"
                placeholder="Enter your last name..."
              />
            </div>
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
            <input
              value={repassword}
              onChange={(e) => setRepassword(e.target.value)}
              type="password"
              name="re-password"
              placeholder="Enter your password again..."
            />

            <button type="submit">Sign Up</button>
          </li>
        </form>
      </ul>
    </div>
  );
}

export default Register;
