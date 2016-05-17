window.Login = React.createClass({
  render: function() {

    return (
    <div>
      <br />
      <form>
        <input type="text" placeholder="Username" /><br /><br />
        <input type="text" placeholder="Password" /><br /><br />
        <button>Login</button>
        <button>Register</button>
      </form>
     </div>
    );
  }
});