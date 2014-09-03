module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    uglify: {
      options: {
        sourceMap: true
      },
      dist: {
        files: {
          'public/javascripts/build.min.js': 'private/javascripts/**/*.js'
        }
      }
    },
    sass: {
      dist: {
        files: [{
          expand: true,
          cwd: 'private/stylesheets',
          src: ['*.scss'],
          dest: 'public/stylesheets',
          ext: '.css'
        }]
      }
    },
    watch: {
      stylesheets: {
        files: ['private/stylesheets/**/*.scss'],
        tasks: ['sass']
      },
      javascripts: {
        files: ['private/javascripts/**/*.js'],
        tasks: ['uglify']
      }
    },
    clean: [
      'public/javascripts/**/*.min.js',
      'public/javascripts/**/*.map',
      'public/stylesheets/**/*.css',
      'public/stylesheets/**/*.css.map',
    ]
  });

  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-clean');

  grunt.registerTask('default', ['uglify', 'sass']);

};