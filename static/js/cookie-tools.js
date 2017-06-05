document.write(" <script language='javascript' src='/js/js.cookie.js' > </script>"); 
function getCookieValue(name) {
	return Cookies.get(name);
}

function getCookieValueEscaped(name) {
	var encodedValue = Cookies.get(name);
    if (encodedValue != null) {
      var value = unescape(encodedValue);
      if (value != null) {
        return value;
      }
    }
    return encodedValue;
}

function setCookieValue(name, value) {
	Cookies.set(name, value);
}

function setCookieValueEscaped(name, value) {
	var encodedValue = escape(value);
        Cookies.set(name, encodedValue);
}