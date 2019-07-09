/*
 * Global js
 */
$(document)
.ready(function() {
  // fix menu when passed
  $('.masthead')
    .visibility({
      once: false,
      onBottomPassed: function() {
        $('.fixed.menu').transition('fade in');
      },
      onBottomPassedReverse: function() {
        $('.fixed.menu').transition('fade out');
      }
    })
  ;

  // create sidebar and attach to menu open
  $('.ui.sidebar')
    .sidebar('attach events', '.toc.item')
  ;

  $('.list-container')
    .transition('slide right in')
  ;

  $('.side-container')
    .transition('slide down in')
  ;
  
  $('.ui.sticky')
    .sticky({
      context: '.content-container'
    })
  ;
});