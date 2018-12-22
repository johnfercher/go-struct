using System;
using System.Threading.Tasks;
using AnyApplication.Domain.Entities;
using AnyApplication.Domain.Enums;

namespace AnyApplication.Domain.Interfaces.Repositories
{
    public interface IDocumentDbRepository<T> where T : ProviderResponse
    {
        List<User> GetUsers();
        Task<bool> AddUser(UserDto user);
    }
}
