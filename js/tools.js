//JS "singleton" pattern
const MtTools = (function() {
  //// PRIVATE
  function some_function() {
  }


  //// PUBLIC
  return {
    /**
    * Creates DOM element.
    */
    Tag: function (tagName, innerHTML, classes, element_id) {
      let element = document.createElement(tagName ? tagName : 'div');

      if (innerHTML) {
        element.innerHTML = innerHTML;
      }

      if (classes) {
        if (!Array.isArray(classes)) {
          classes = [classes];
        }

        element.classList.add(...classes);
      }

      if (element_id) {
        element.id = element_id;
      }

      return element;
    },

    /**
    * Creates Font-Awesome icon element.
    */
    Icon: function (iconName, title, classes) {
      let icon_classes = ['far', 'fa-' + iconName];

      if(classes)
      {
        if(Array.isArray(classes))
        {
          icon_classes.push(...classes);
        }
        else
        {
          icon_classes.push(classes);
        }
      }

      let icon_element = this.Tag('i', null, icon_classes);

      if(title)
      {
        icon_element.title = title;
      }

      return icon_element;
    },

    //#region AJAX helpers
    /**
    * Fetch URL (with GET request) and replace element's contents.
    * @param element HTMLElement
    * @param url string
    */
    SimpleGet: function (element, url, placeholder = null, success_callback = null) {
      if (placeholder == null) {
        placeholder = this.Icon('spinner', 'waiting', 'fa-pulse').outerHTML + ' Sending request...';
      }

      element.innerHTML = placeholder;

      fetch(url)
        .then(function (response) {
          if (response.ok) {
            return response.text(); //reads response as HTML-string and returns next promise
          } else {
            console.error(response);
            return null;
          }
        })
        .then(function (html) {
          element.innerHTML = html;

          if (typeof success_callback === 'function') {
            success_callback();
          }
        })
        .catch(error => console.error(error));
    },
    //#endregion
  };
})();
