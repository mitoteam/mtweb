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
  };
})();
