using System;
using System.Threading.Tasks;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using AnyApplication.Domain.Dtos;
using AnyApplication.Domain.Interfaces;
using AnyApplication.Domain.Interfaces.Services;
using AnyApplication.Filters;
using AnyApplication.Models;

namespace AnyApplication.Controllers
{
    [Route("api/[controller]/v2")]
    public class UserController : Controller
    {
        public readonly IUserRepository UserRepository;
        public ILogger Logger { get; private set; }
        public IPublisherUser PublisherUser { get; }

        public UserController(IUserRepository userRepository, ILogger logger)
        {
            UserRepository = userRepository;
            Logger = logger ?? throw new ArgumentNullException(nameof(logger));
            PublisherUser = new PublisherUser();
        }

        [HttpGet]
        public IActionResult GetUsers()
        {
            return Ok(UserRepository.GetUsers());
        }

        [HttpPost]
        [ServiceFilter(typeof(FluentValidationFilter))]
        public async Task<IActionResult> AddUser([FromBody]User user)
        {
            var userDto = Mapper.Map<UserDto>(user);
            var result = await UserRepository.AddUser(userDto);

            if (!result)
                return InternalServerError();

            return Ok();        
        }
    }
}