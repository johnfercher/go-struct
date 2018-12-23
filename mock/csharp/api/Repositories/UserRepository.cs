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
    public class UserRepository : IUserRepository
    {
        public List<User> GetUsers()
        {
            return new List<User>();
        }

        public bool AddUser(UserDto user)
        {
            return true;
        }

        public async Task<bool> AddUserAsync(UserDto user)
        {
            return await Task.FromResult(true);
        }
    }
}
