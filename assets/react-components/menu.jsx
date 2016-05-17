window.Menu = React.createClass({
  render: function() {

    return (
      <nav className='navbar'>
      <div className='container'>
      	<ul className='navbar-list'>
    			<li className='navbar-item'><a className='navbar-link' href='/'>Home</a></li>
    			<li className='navbar-item'><a className='navbar-link' href='/about'>About Us</a></li>
      	</ul>
        </div>
      </nav>
    );
  }
});